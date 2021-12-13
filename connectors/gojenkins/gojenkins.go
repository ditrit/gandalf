package gojenkins

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gandalf/connectors/gogithub/client"
	"gandalf/connectors/gogithub/poll"
	"os"

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
		worker.RegisterCommandsFuncs("CREATE_ISSUE", CreateIssue)
		worker.RegisterCommandsFuncs("CREATE_PULL", CreatePull)

		scanService := new(poll.ScanService)
		worker.RegisterServicesFuncs("ScanService", scanService.Start)

		worker.Run()
	}
}
