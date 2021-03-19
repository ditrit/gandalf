package poll

import (
	"gandalf/connectors/gogithub/pull"
	"gandalf/connectors/gogithub/repository"
	"time"

	"github.com/google/go-github/v33/github"
)

type ScanService struct {
	LastCommit *time.Time
	LastPull   *time.Time
}

func (ss ScanService) Scan(clientGandalf *goclient.ClientGandalf, client *github.Client, actions []string, owner, repo string) {
	for range time.Tick(time.Minute * 1) {
		go ss.poll(client, actions, owner, repo)
	}

}

func (ss ScanService) poll(clientGandalf *goclient.ClientGandalf, client *github.Client, actions []string, owner, repo string) {
	for _, action := range actions {
		if action == "commit" {
			commit := repository.GetLastCommitsRepository(client, owner, repo)
			if commit.Commit.Committer.Date.After(ss.LastCommit) {
				ss.LastCommit = commit.Commit.Committer.Date
				//EVENT
				clientGandalf.SendEvent(topic, event string, options map[string]string)
			}
		} else if action == "pull" {
			pull := pull.GetLastPullRequest(client, owner, repo)
			if pull.MergedAt.After(ss.LastPull) {
				ss.LastPull = pull.MergedAt
				//EVENT
				clientGandalf.SendEvent(topic, event string, options map[string]string)
			}
		}
	}
}
