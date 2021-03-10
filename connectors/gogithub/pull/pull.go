package pull

import (
	"context"
	"log"

	"github.com/google/go-github/v33/github"
)

func ListPullRequest(client *github.Client, owner, repo string) []*github.PullRequest {
	ctx := context.Background()
	pulls, _, err := client.Pulls.List(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	return pulls
}
