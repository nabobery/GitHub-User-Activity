# GitHub User Activity CLI

https://roadmap.sh/projects/github-user-activity

### Overview

The **GitHub Activity CLI** is a command-line interface application written in Go. It allows users to fetch and display the recent activities of a GitHub user.

---

### Features

- **Fetch User Activities**: Retrieves the latest activities of a specified GitHub user.
- **Display Activity Summary**: Outputs a summary of the user's activities, including pushes, stars, pull requests, and issue comments.

---

### Prerequisites

- [Go Programming Language](https://golang.org/) installed (version 1.16 or later).

---

## Requirements

- Go programming language installed on your machine.
- Basic understanding of command line operations.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/nabobery/github-user-activity.git
   cd github-user-activity
   ```

2. Build the application:

   ```bash
   # For Linux/Mac
   go build -o github-activity

   # For Windows
   go build -o github-activity.exe main.go
   ```

### Usage

To fetch and display the activity of a GitHub user, use the following command:

```bash
github-activity <username>
```

Replace `<username>` with the GitHub username you want to check.

Example:

```bash
github-activity 3b1b
```

Example Output:

```bash
- Pushed 1 commit(s) to 3b1b/3Blue1Brown.com
- Closed a pull request in 3b1b/3Blue1Brown.com
- Pushed 1 commit(s) to 3b1b/3Blue1Brown.com
- Closed a pull request in 3b1b/3Blue1Brown.com
- Pushed 1 commit(s) to 3b1b/videos
- Added a comment to an issue in 3b1b/manim
- Added a comment to an issue in 3b1b/manim
- Pushed 1 commit(s) to 3b1b/manim
- Closed a pull request in 3b1b/manim
- Opened a pull request in 3b1b/manim
- Pushed 1 commit(s) to 3b1b/manim
- Pushed 1 commit(s) to 3b1b/3Blue1Brown.com
- Pushed 1 commit(s) to 3b1b/3Blue1Brown.com
- Closed a pull request in 3b1b/3Blue1Brown.com
- Pushed 1 commit(s) to JustinGamer191/3Blue1Brown.com
```

---

### How It Works

- **GitHub API**: The CLI uses the GitHub API to fetch user activity data.
- **Cobra Library**: The CLI is powered by [Cobra](https://github.com/spf13/cobra), a library for building powerful Go CLIs.
- **Activity Types**: The CLI processes different activity types, including `PushEvent`, `WatchEvent`, `PullRequestEvent`, and `IssueCommentEvent`.

---

## Extending the Project

1. **Data Persistence**:
   - Save fetched activity data to a file or database.
   - Implement caching to avoid redundant API calls.

2. **Additional Features**:
   - Add filtering options for specific activity types or repositories.
   - Implement pagination for handling large activity lists.

---

## License

This project is open-source and free to use under the [MIT License](LICENSE). Contributions are welcome!
