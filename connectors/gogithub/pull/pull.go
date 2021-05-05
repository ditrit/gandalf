package pull

import (
	"context"
	"log"

	"github.com/google/go-github/v33/github"
)

type ListPullRequestPayload struct {
	Owner      string
	Repository string
}

func ListPullRequest(client *github.Client, owner, repo string) []*github.PullRequest {
	ctx := context.Background()
	pulls, _, err := client.PullRequests.List(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	return pulls
}

type GetLastPullRequestPayload struct {
	Owner      string
	Repository string
}

func GetLastPullRequest(client *github.Client, owner, repo string) *github.PullRequest {
	pulls := ListPullRequest(client, owner, repo)
	var lastPull *github.PullRequest
	for _, pull := range pulls {
		if lastPull == nil {
			lastPull = pull
		} else {
			//if pull.MergedAt > lastPull.MergedAt {
			if pull.MergedAt.After(*lastPull.MergedAt) {
				lastPull = pull
			}
		}
	}
	return lastPull
}
