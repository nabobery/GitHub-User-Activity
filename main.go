package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "github-activity [username]",
	Short: "Get activities of a github user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		formattedUrl := fmt.Sprintf("https://api.github.com/users/%s/events", username)
		resp, err := http.Get(formattedUrl)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK { // Check for non-200 status codes
			fmt.Printf("Error: received status code %d\n", resp.StatusCode)
			return
		}

		activityData, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading body:", err)
			return
		}

		var activities []map[string]interface{}
		if err := json.Unmarshal(activityData, &activities); err != nil {
			log.Fatalf("Error unmarshalling activity data: %v\n", err)
			os.Exit(1)
		}

		for _, activity := range activities {
			activityType, ok := activity["type"].(string)
			if !ok {
				log.Println("Skipping activity with invalid type")
				continue
			}

			repo, ok := activity["repo"].(map[string]interface{})
			if !ok {
				log.Println("Skipping activity with invalid repo")
				continue
			}
			repoName, ok := repo["name"].(string)
			if !ok {
				log.Println("Skipping activity with invalid repo name")
				continue
			}

			switch activityType {
			case "PushEvent":
				payload, ok := activity["payload"].(map[string]interface{})
				if !ok {
					log.Println("Skipping push event with invalid payload")
					continue
				}
				commits, ok := payload["commits"].([]interface{})
				if !ok {
					log.Println("Skipping push event with invalid commits")
					continue
				}
				fmt.Printf("- Pushed %d commit(s) to %s\n", len(commits), repoName)
			case "WatchEvent":
				fmt.Printf("- Starred %s\n", repoName)
			case "PullRequestEvent":
				payload, ok := activity["payload"].(map[string]interface{})
				if !ok {
					log.Println("Skipping pull request event with invalid payload")
					continue
				}
				action, ok := payload["action"].(string)
				if !ok {
					log.Println("Skipping pull request event with invalid action")
					continue
				}
				if action == "closed" {
					fmt.Printf("- Closed a pull request in %s\n", repoName)
				} else if action == "opened" {
					fmt.Printf("- Opened a pull request in %s\n", repoName)
				}
			case "IssueCommentEvent":
				payload, ok := activity["payload"].(map[string]interface{})
				if !ok {
					log.Println("Skipping issue comment event with invalid payload")
					continue
				}
				action, ok := payload["action"].(string)
				if !ok {
					log.Println("Skipping issue comment event with invalid action")
					continue
				}
				if action == "created" {
					fmt.Printf("- Added a comment to an issue in %s\n", repoName)
				}
			}
		}
	},
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
