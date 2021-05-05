package poll

import (
	"fmt"
	"strings"
	"time"

	"github.com/ditrit/gandalf/libraries/goclient"

	"github.com/ditrit/gandalf/connectors/gogithub/repository"

	"github.com/ditrit/gandalf/connectors/gogithub/pull"

	"github.com/ditrit/gandalf/core/models"

	"github.com/google/go-github/v33/github"
)

type ScanService struct {
	LastCommit time.Time
	LastPull   time.Time
}

func (ss ScanService) Start(context map[string]interface{}, clientGandalf *goclient.ClientGandalf) {
	clientGithub, ok := context["client"].(*github.Client)
	if ok {
		eventTypeToPolls, ok := context["EventTypeToPolls"].([]models.EventTypeToPoll)
		fmt.Println("eventTypeToPolls")
		fmt.Println(eventTypeToPolls)
		if ok {
			for range time.Tick(time.Minute * 1) {
				fmt.Println("POLL")
				for _, eventTypeToPoll := range eventTypeToPolls {
					if eventTypeToPoll.EventType.Name == "COMMIT" {
						fmt.Println("COMMIT")
						resourceSplit := strings.Split(eventTypeToPoll.Resource.Name, ":")
						commit := repository.GetLastCommitsRepository(clientGithub, resourceSplit[0], resourceSplit[1])
						fmt.Println("commit")
						fmt.Println(commit)
						fmt.Println("TEST")
						fmt.Println(commit.Commit.Committer.Date)
						fmt.Println(ss.LastCommit)
						fmt.Println(commit.Commit.Committer.Date.After(ss.LastCommit))
						fmt.Println("END TEST")
						if commit.Commit.Committer.Date.After(ss.LastCommit) {
							fmt.Println("after")
							ss.LastCommit = *commit.Commit.Committer.Date
							//EVENT
							fmt.Println("SEND EVENT COMMIT")
							clientGandalf.SendEvent(eventTypeToPoll.Resource.Name, eventTypeToPoll.EventType.Name, nil)
						} else {
							ss.LastCommit = *commit.Commit.Committer.Date
						}
					} else if eventTypeToPoll.EventType.Name == "PULL" {
						fmt.Println("PULL")
						resourceSplit := strings.Split(eventTypeToPoll.Resource.Name, ":")
						pull := pull.GetLastPullRequest(clientGithub, resourceSplit[0], resourceSplit[1])
						if pull.MergedAt.After(ss.LastPull) {
							ss.LastPull = *pull.MergedAt
							//EVENT
							fmt.Println("SEND EVENT PULL")
							clientGandalf.SendEvent(eventTypeToPoll.Resource.Name, eventTypeToPoll.EventType.Name, nil)
						} else {
							ss.LastPull = *pull.MergedAt
						}
					}
				}
			}
		}
	}
}
