package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v33/github"
)

func CreateRepository(client *github.Client, name, description string, private bool) *github.Repository {
	ctx := context.Background()
	r := &github.Repository{Name: name, Private: private, Description: description}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
	return repo
}

func CreateRepositoryFromTemplate(client *github.Client, templateOwner, templateRepo, name, owner, description string, private bool) *github.Repository {
	ctx := context.Background()
	r := &github.TemplateRepoRequest{Name: name, Owner: owner, Private: private, Description: description}
	repo, _, err := client.Repositories.CreateFromTemplate(ctx, templateOwner, templateRepo, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created")
	return repo
}

func DeleteRepository(client *github.Client, owner, repo string) {
	ctx := context.Background()
	_, err := client.Repositories.Delete(ctx, owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully delete")
}

func ListCommitsRepository(client *github.Client, owner, repo string) []*github.RepositoryCommit {
	ctx := context.Background()
	commits, _, err := client.ListCommits(ctx, owner, repo, nil)
	if err != nil {
		log.Fatal(err)
	}
	return commits
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
