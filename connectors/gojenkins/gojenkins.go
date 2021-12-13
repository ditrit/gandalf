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
	if err == nil {
		if inputPayload.Username != "" && inputPayload.Password != "" {
			clientJenkins := client.ClientWithAuthentication(inputPayload.URL, inputPayload.Username, inputPayload.Password)
			worker.Context["client"] = clientJenkins
		} else if inputPayload.Username != "" && inputPayload.Password != "" {
			fmt.Println("BasicAuthentification")
			clientJenkins := client.ClientWithoutAuthentication(inputPayload.URL)
			worker.Context["client"] = clientJenkins
		}

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

type InputPayload struct {
	Username string
	Password string
	URL      string
	//....
}
