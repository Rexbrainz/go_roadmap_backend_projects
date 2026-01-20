package activity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Events struct {
	Type string
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload	json.RawMessage `json:"payload"`
}

func FormatPushEvent(push json.RawMessage, repo string) (string, error) {
	var p struct {
		Ref string`json:"ref"`
		Before string `json:"before"`
		Head string `json:"head"`
	}

	if err := json.Unmarshal(push, &p); err != nil {
		return "", err
	}
	return fmt.Sprintf("Pushed updates to %s", repo), nil
}

func FormatIssueEvent(push json.RawMessage, repo string) (string, error){
	var p struct {
		Action string `json:"action"`
	}

	if err := json.Unmarshal(push, &p); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s an issue in %s", p.Action, repo), nil
}

func FormatWatchEvent(push json.RawMessage, repo string) (string, error) {
	var p struct {
		Action string `json:"action"`
	}

	if err := json.Unmarshal(push, &p); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", p.Action, repo), nil
}

func FormatForkEvent(push json.RawMessage, repo string) (string, error) {
	var p struct {
		Forkee string `json:"forkee"`
	}

	if err := json.Unmarshal(push, &p); err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s a %s", p.Forkee, repo), nil
}

func FormatPullRequestEvent(push json.RawMessage, repo string) (string, error) {
	var p struct {
		Action string `json:"action"`
	}

	if err := json.Unmarshal(push, &p); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s pull request in %s", p.Action, repo), nil
}

func FormatEvents(event Events) (string, error) {
	switch event.Type {
	case "PushEvent":
		return FormatPushEvent(event.Payload, event.Repo.Name)
	case "IssuesEvent":
		return FormatIssueEvent(event.Payload, event.Repo.Name)
	case "WatchEvent":
		return FormatWatchEvent(event.Payload, event.Repo.Name)
	case "ForkEvent":
		return FormatForkEvent(event.Payload, event.Repo.Name)
	case "PullRequestEvent":
		return FormatPullRequestEvent(event.Payload, event.Repo.Name)

	default:
		return fmt.Sprintf("Perfomed %s on %s", event.Type, event.Repo.Name), nil
	}
}

func HandlePassedEventType(events []Events, eventType string) {
	found := false
	for _, event := range events {
		if event.Type == eventType {
			activity, err := FormatEvents(event)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(activity)
			found = true
		}
	}
	if !found {
		fmt.Printf("%s is not included in the latest user activities", eventType)
	}
}

func ParseAndListActivities(data []byte, filter *string) {
	var events []Events

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&events)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 
	}

	if filter != nil {
		HandlePassedEventType(events, *filter)
		return
	}
	
	for _, event := range events {
		activity, err := FormatEvents(event)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(activity)
	}
}
