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

Github user activity lists by default the Recent activities
of the provided username on the standard output.
Flags:
	-t: Specifies an exact event type to list.`

type Client struct {
	BaseURL    string
	ApiKey     string
	HTTPClient *http.Client
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

func HandleResponseStatusError(status int) error {
	switch status {
	case http.StatusOK:
	case http.StatusForbidden:
		return fmt.Errorf("Github API rate limit exceeded")
	case http.StatusNotFound:
		return fmt.Errorf("User not found")
	}
	return nil
}

func (c *Client) GetActivities(URL string) ([]byte, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	if c.ApiKey != "" {
		req.Header.Set("Authorization", "Bearer " + c.ApiKey)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if err := HandleResponseStatusError(resp.StatusCode); err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *eventType != "" {
		ParseAndListActivities(activities, eventType)
		return
	}
	ParseAndListActivities(activities, nil)
}
