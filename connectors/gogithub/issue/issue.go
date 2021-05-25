package issue

import (
	"context"

	"github.com/google/go-github/v33/github"
)

type CreateIssuePayload struct {
	Token      string
	Owner      string
	Repository string
	Title      string
	Body       string
}

func CreateIssue(client *github.Client, owner, repository, title, body string) (err error) {
	ctx := context.Background()
	r := &github.IssueRequest{Title: &title, Body: &body}
	_, _, err = client.Issues.Create(ctx, owner, repository, r)

	return
}
