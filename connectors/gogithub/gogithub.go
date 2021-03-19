package main

import (
	"bufio"
	"fmt"
	"gandalf/libraries/goclient"
	"os"
	"shoset/msg"

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

	//GET VALUE FROM STDIN
	var inputPayload InputPayload
	err := json.Unmarshal([]byte(input.Text()), &inputPayload)
	//CREATE AUTHENTIFICATION
	clientGithub := client.BasicAuthentification(inputPayload.Username, inputPayload.Password)
	worker.Context["client"] = clientGithub


	worker.RegisterCommandsFuncs("CREATE_REPOSITORY", CreateRepository)
	worker.RegisterCommandsFuncs("CREATE_REPOSITORY_FROM_TEMPLATE", CreateRepositoryFromTemplate)
	worker.RegisterCommandsFuncs("DELETE_REPOSITORY", DeleteRepository)

	worker.Run()
}

type InputPayload struct {
	Username string
	Password string
	//ETC....
}

func CreateRepository(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var createRepositoryPayload repository.CreateRepositoryPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &createRepositoryPayload)
	if err == nil {
		clientGithub, ok := worker.Context["client"].(*github.Client)
		if ok {
			repository.CreateRepository(clientGithub, createRepositoryPayload.Name, createRepositoryPayload.Description, createRepositoryPayload.Private)
		}
	}
}

func CreateRepositoryFromTemplate(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var createRepositoryFromTemplatePayload repository.CreateRepositoryFromTemplatePayload
	err := json.Unmarshal([]byte(command.GetPayload()), &createRepositoryFromTemplatePayload)
	if err == nil {
		clientGithub, ok := worker.Context["client"].(*github.Client)
		if ok {
			repository.CreateRepositoryFromTemplate(clientGithub, createRepositoryFromTemplatePayload.TemplateOwner, createRepositoryFromTemplatePayload.TemplateRepo, createRepositoryFromTemplatePayload.Name, createRepositoryFromTemplatePayload.Owner, createRepositoryFromTemplatePayload.Description, createRepositoryFromTemplatePayload.Private)
		}
	}
}

func DeleteRepository(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	var deleteRepositoryPayload repository.DeleteRepositoryPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &deleteRepositoryPayload)
	if err == nil {
		clientGithub, ok := worker.Context["client"].(*github.Client)
		if ok {
			repository.CreateRepositoryFromTemplate(clientGithub, deleteRepositoryPayload.Owner, deleteRepositoryPayload.Repository)
		}
	}
}



func Scan(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	
}
