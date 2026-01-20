# GitHub User Activity CLI (Go)

A command-line application written in Go that fetches and displays the **recent public GitHub activity** of a user using the GitHub Events API.

This project is based on the **GitHub User Activity** project from the :contentReference[oaicite:0]{index=0} backend projects list.

🔗 Project reference:(https://roadmap.sh/projects/github-user-activity)

---

## 🚀 Overview

The GitHub User Activity CLI allows you to inspect a user's recent public GitHub activity directly from the terminal.  
It consumes GitHub’s **public events API**, parses event-specific payloads, and presents the information in a human-readable format.

The project focuses on:
- API consumption
- JSON parsing
- CLI design
- Defensive backend programming
- Testable code structure

---

## ✨ Features

- Fetches recent **public GitHub events** for a given username
- Formats common event types into readable messages
- Supports filtering by **event type**
- Handles common API errors gracefully
- Uses **only the Go standard library**

---

## 🛠 Tech Stack

- Go
- `net/http`
- `encoding/json`
- GitHub REST API

---

## 📦 Supported Event Types

The CLI provides custom formatting for common GitHub events:

- `PushEvent`
- `IssuesEvent`
- `WatchEvent` (starred repositories)
- `ForkEvent`
- `PullRequestEvent`

All other event types fall back to a generic formatter.

### Important Note on Push Events

The GitHub endpoint used (`/users/:username/events`) does **not** expose commit counts for `PushEvent`.

As a result:
- Push events are displayed as **repository updates**
- Commit totals are intentionally not shown

This behavior reflects the **actual data returned by the API**, not an implementation limitation.

---

## 🧪 Error Handling

The application handles the following error cases:

- **Invalid username** → `404 Not Found`
- **GitHub API rate limit exceeded** → `403 Forbidden`

Authentication is **optional**.  
If the environment variable `GITHUB_FGT_API_TOKEN` is set, it will be used to increase API rate limits.  
The application still works correctly without a token.

---

## 🚀 Usage

### Run the application
```bash
./program_name [OPTION] <username>
OPTION filters by event type.
./program_name -t Event <username>
