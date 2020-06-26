package client

import (
	"fmt"
	"log"

	gitlab "github.com/xanzy/go-gitlab"
)

type ClientGitlab struct {
	Token  string
	Client *gitlab.Client
}

func NewClient(token string) *ClientGitlab {
	client := new(ClientGitlab)
	client.Token = token

	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	client.Client = git

	return client

}

func NewClientBasic(login, password, rooturl string) *ClientGitlab {
	client := new(ClientGitlab)
	git, err := gitlab.NewBasicAuthClient(
		login,
		password,
		gitlab.WithBaseURL(rooturl),
	)
	if err != nil {
		log.Fatal(err)
	}
	client.Client = git
	fmt.Println(git)
	return client
}
