package main

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/ditrit/gandalf/connectors/gogithub/client"
	"github.com/ditrit/gandalf/connectors/gogithub/poll"
	"github.com/ditrit/gandalf/connectors/gogithub/repository"
	"github.com/ditrit/gandalf/libraries/goclient"

	"github.com/ditrit/gandalf/core/models"

	"os"

	"github.com/ditrit/shoset/msg"

	"github.com/google/go-github/v33/github"

	worker "github.com/ditrit/gandalf/connectors/go"
)

//main : main
func main() {

	var major = int64(1)
	var minor = int64(0)

	fmt.Println("VERSION")
	fmt.Println(major)
	fmt.Println(minor)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())

	worker := worker.NewWorker(major, minor)

	var inputPayload InputPayload
	err := json.Unmarshal([]byte(input.Text()), &inputPayload)
	fmt.Println("err")
	fmt.Println(err)
	fmt.Println("InputPayload")
	fmt.Println(inputPayload)
	fmt.Println(inputPayload.EventTypeToPolls)
	if err == nil {
		if inputPayload.Token != "" {
			fmt.Println("Oauth2Token")
			clientGithub := client.Oauth2Authentification(inputPayload.Token)
			worker.Context["client"] = clientGithub
		} else if inputPayload.Username != "" && inputPayload.Password != "" {
			fmt.Println("BasicAuthentification")
			clientGithub := client.BasicAuthentification(inputPayload.Username, inputPayload.Password)
			worker.Context["client"] = clientGithub
		}

		worker.Context["EventTypeToPolls"] = inputPayload.EventTypeToPolls

		worker.RegisterCommandsFuncs("CREATE_REPOSITORY", CreateRepository)
		worker.RegisterCommandsFuncs("CREATE_REPOSITORY_FROM_TEMPLATE", CreateRepositoryFromTemplate)
		worker.RegisterCommandsFuncs("DELETE_REPOSITORY", DeleteRepository)

		scanService := new(poll.ScanService)
		worker.RegisterServicesFuncs("ScanService", scanService.Start)

		worker.Run()
	}
}

type InputPayload struct {
	Username         string
	Password         string
	Token            string
	EventTypeToPolls []models.EventTypeToPoll
	//....
}

func CreateRepository(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var createRepositoryPayload repository.CreateRepositoryPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &createRepositoryPayload)
	if err == nil {
		var clientGithub *github.Client
		if createRepositoryPayload.Token != "" {
			clientGithub = client.Oauth2Authentification(createRepositoryPayload.Token)
		} else if createRepositoryPayload.Username != "" && createRepositoryPayload.Password != "" {
			clientGithub = client.BasicAuthentification(createRepositoryPayload.Username, createRepositoryPayload.Password)
		} else {
			clientGithub, _ = context["client"].(*github.Client)
		}

		err = repository.CreateRepository(clientGithub, createRepositoryPayload.Name, createRepositoryPayload.Description, createRepositoryPayload.Private)
		if err == nil {
			return 0
		}
	}
	return 1
}

func CreateRepositoryFromTemplate(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var createRepositoryFromTemplatePayload repository.CreateRepositoryFromTemplatePayload
	err := json.Unmarshal([]byte(command.GetPayload()), &createRepositoryFromTemplatePayload)
	if err == nil {
		var clientGithub *github.Client
		if createRepositoryFromTemplatePayload.Token != "" {
			clientGithub = client.Oauth2Authentification(createRepositoryFromTemplatePayload.Token)
		} else if createRepositoryFromTemplatePayload.Username != "" && createRepositoryFromTemplatePayload.Password != "" {
			clientGithub = client.BasicAuthentification(createRepositoryFromTemplatePayload.Username, createRepositoryFromTemplatePayload.Password)
		} else {
			clientGithub, _ = context["client"].(*github.Client)
		}
		err = repository.CreateRepositoryFromTemplate(clientGithub, createRepositoryFromTemplatePayload.TemplateOwner, createRepositoryFromTemplatePayload.TemplateRepo, createRepositoryFromTemplatePayload.Name, createRepositoryFromTemplatePayload.Owner, createRepositoryFromTemplatePayload.Description, createRepositoryFromTemplatePayload.Private)
		if err == nil {
			return 0
		}
	}
	return 1
}

func DeleteRepository(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var deleteRepositoryPayload repository.DeleteRepositoryPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &deleteRepositoryPayload)
	if err == nil {
		var clientGithub *github.Client
		if deleteRepositoryPayload.Token != "" {
			clientGithub = client.Oauth2Authentification(deleteRepositoryPayload.Token)
		} else if deleteRepositoryPayload.Username != "" && deleteRepositoryPayload.Password != "" {
			clientGithub = client.BasicAuthentification(deleteRepositoryPayload.Username, deleteRepositoryPayload.Password)
		} else {
			clientGithub, _ = context["client"].(*github.Client)
		}

		err = repository.DeleteRepository(clientGithub, deleteRepositoryPayload.Owner, deleteRepositoryPayload.Repository)
		if err == nil {
			return 0
		}
	}
	return 1
}
