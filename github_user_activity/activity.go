package activity

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const Usage = `Usage: github-activity [OPTION] Username

			The program lists by default the Recent activities
			of the provided username on the standard output.
			Flags:
				-t: Specifies an exact event type to list.`

type Client struct {
	BaseURL    string
	ApiKey     string
	HTTPClient *http.Client
}

type Activities struct {
	Events []struct {
		Type    string
		Repo    []string
		Payload []string
	}
}

type Events struct {
	Type string
	Repo struct {
		Name string
	}
}

func NewClient() *Client {
	return &Client{
		BaseURL: "https://api.github.com",
		ApiKey:  os.Getenv("GITHUB_FGT_API_TOKEN"),
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) FormatURL(username string) string {
	return fmt.Sprintf("%s/users/%s/events", c.BaseURL, username)
}

func (c *Client) GetActivities(URL string) ([]byte, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	if c.ApiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected response status %q", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ParseAndPrintActivities(data []byte) error {
	return nil
}

func ParseAndPrintActivity(data []byte, event string) error {
	return nil
}

func Main() {
	eventType := flag.String("t", "", "Gets activities based on event type")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println(Usage)
		os.Exit(0)
	}

	args := flag.Args()
	c := NewClient()
	URL := c.FormatURL(args[0])
	activities, err := c.GetActivities(URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %v", err)
		os.Exit(1)
	}

	if eventType != "" {
		if err := ParseAndPrintActivity(activities, eventType); err != nil {
			fmt.Fprintf(os.Stderr, "unexpected error: %v", err)
			os.Exit(1)
		}
		return
	}
	if err := ParseAndPrintActivities(activities); err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %v", err)
		os.Exit(1)
	}
}
