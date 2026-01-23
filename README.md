# Go Backend Projects

Monorepo of backend projects implemented in **Go** following the roadmap.sh backend track.  
Each project lives in its own folder at the repo root and ships with its own docs.

🔗 roadmap.sh projects: https://roadmap.sh/backend/projects

---

## 🚀 Vision
- Practice backend engineering fundamentals with Go
- Build a portfolio of real, runnable projects
- Demonstrate practical skills for backend/cloud roles
- Learn by doing through the roadmap.sh project list

---

## 📦 Projects Overview

| Project | Status | Description | Link |
|--------|--------|-------------|------|
| **Task Tracker** | ✅ Completed | CLI to add/update/delete tasks with local JSON persistence | https://roadmap.sh/projects/task-tracker |
| **GitHub User Activity CLI** | ✅ Completed | CLI that fetches and formats a user's recent GitHub public events | https://roadmap.sh/projects/github-user-activity |
| **Weather API** | ✅ Completed | HTTP API that proxies OpenWeatherMap with in-memory/Redis caching | https://roadmap.sh/projects/weather-api-wrapper-service |

Projects will be added here as they are completed.

---

## 🗂 Repository Structure

```text
go-backend-projects/
├── README.md                # Monorepo overview (this file)
├── task-tracker/            # Task Tracker CLI project
├── github_user_activity/    # GitHub User Activity CLI project
└── weather_api/             # Weather API project
```

---

## 🧪 Running Tests
Each project is its own Go module. Run tests from within the project directory:

```bash
cd task-tracker && go test ./...
cd github_user_activity && go test ./...
cd weather_api && go test ./...
```

---

## 🛠 Prerequisites
- Go toolchain (modules target Go 1.25.x as declared in go.mod files)
- For `weather_api`: an OpenWeatherMap API key (see `weather_api/README.md` for setup)

