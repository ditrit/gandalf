package pull

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v33/github"
)

func List(client *github.Client, owner, repo string) {
	ctx := context.Background()
	pulls, _, err := client.Pulls.List(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully listed")
}
