package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v33/github"
)

type CreateRepositoryPayload struct {
	Username    string
	Password    string
	Token       string
	Name        string
	Description string
	Private     bool
}

func CreateRepository(client *github.Client, name, description string, private bool) (err error) {
	ctx := context.Background()
	r := &github.Repository{Name: name, Private: private, Description: description}
	_, _, err = client.Repositories.Create(ctx, "", r)

	return
}

type CreateRepositoryFromTemplatePayload struct {
	Username      string
	Password      string
	Token         string
	TemplateOwner string
	TemplateRepo  string
	Name          string
	Owner         string
	Description   string
	Private       bool
}

func CreateRepositoryFromTemplate(client *github.Client, templateOwner, templateRepo, name, owner, description string, private bool) (err error) {
	ctx := context.Background()
	r := &github.TemplateRepoRequest{Name: name, Owner: owner, Private: private, Description: description}
	_, _, err = client.Repositories.CreateFromTemplate(ctx, templateOwner, templateRepo, r)

	return
}

type DeleteRepositoryPayload struct {
	Username   string
	Password   string
	Token      string
	Owner      string
	Repository string
}

func DeleteRepository(client *github.Client, owner, repo string) (err error) {
	ctx := context.Background()
	_, err = client.Repositories.Delete(ctx, owner, repo)

	return
}

type ListCommitsRepositoryPayload struct {
	Owner      string
	Repository string
}

func ListCommitsRepository(client *github.Client, owner, repo string) []*github.RepositoryCommit {
	ctx := context.Background()
	commits, _, err := client.ListCommits(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}
	return commits
}

type GetLastCommitsRepositoryPayload struct {
	Owner      string
	Repository string
}

func GetLastCommitsRepository(client *github.Client, owner, repo string) *github.RepositoryCommit {
	commits := ListCommitsRepository(client, owner, repo)
	var lastCommit *github.RepositoryCommit
	for commit := range commits {
		if lastCommit == nil {
			lastCommit == commit
		} else {
			if commit.Commit.Committer.Date > lastCommit.Commit.Committer.Date {
				lastCommit = commit
			}
		}
	}
	return lastCommit
}

type CreateHookRepositoryPayload struct {
	Owner      string
	Repository string
	Config     map[string]interface{}
	Events     []string
	Active     bool
}

func CreateHookRepository(client *github.Client, owner, repo string, config map[string]interface{}, events []string, active bool) {
	ctx := context.Background()
	h := &github.Repository{Config: config, Events: events, Active: active}
	hook, _, err := client.CreateHook(ctx, owner, repo, h)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", hook.Name())
}