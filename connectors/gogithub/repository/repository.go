package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v33/github"
)

func CreateRepository(client *github.Client, name, description string, private bool) {
	ctx := context.Background()
	r := &github.Repository{Name: name, Private: private, Description: description}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}

func CreateRepositoryFromTemplate(client *github.Client, templateOwner, templateRepo, name, owner, description string, private bool) {
	ctx := context.Background()
	r := &github.TemplateRepoRequest{Name: name, Owner: owner, Private: private, Description: description}
	_, err := client.Repositories.CreateFromTemplate(ctx, templateOwner, templateRepo, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created")
}

func DeleteRepository(client *github.Client, owner, repo string) {
	ctx := context.Background()
	_, err := client.Repositories.Delete(ctx, owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully delete")
}
