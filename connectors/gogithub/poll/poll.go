package poll

import (
	"gandalf/connectors/gogithub/pull"
	"gandalf/connectors/gogithub/repository"
	"gandalf/core/models"
	"gandalf/libraries/goclient"
	"strings"
	"time"

	"github.com/google/go-github/v33/github"
)

type ScanService struct {
	LastCommit *time.Time
	LastPull   *time.Time
}

func (ss ScanService) Start(context map[string]interface{}, clientGandalf *goclient.ClientGandalf) {

	clientGithub, ok := context["client"].(*github.Client)
	if ok {
		eventTypeToPolls, ok := context["EventTypeToPolls"].([]models.EventTypeToPoll)
		if ok {
			for range time.Tick(time.Minute * 1) {
				for _, eventTypeToPoll := range eventTypeToPolls {
					if eventTypeToPoll.EventType.Name == "commit" {
						resourceSplit := strings.Split(eventTypeToPoll.Resource.Name, "/")
						commit := repository.GetLastCommitsRepository(clientGithub, resourceSplit[0], resourceSplit[1])
						if commit.Commit.Committer.Date.After(ss.LastCommit) {
							ss.LastCommit = commit.Commit.Committer.Date
							//EVENT
							clientGandalf.SendEvent(topic, event)
						}
					} else if eventTypeToPoll.EventType.Name == "pull" {
						resourceSplit := strings.Split(eventTypeToPoll.Resource.Name, "/")
						pull := pull.GetLastPullRequest(clientGithub, resourceSplit[0], resourceSplit[1])
						if pull.MergedAt.After(ss.LastPull) {
							ss.LastPull = pull.MergedAt
							//EVENT
							clientGandalf.SendEvent(topic, event)
						}
					}
				}
			}
		}
	}
}
