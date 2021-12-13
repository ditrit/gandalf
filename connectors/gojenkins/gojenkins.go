package gojenkins

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gandalf/connectors/gogithub/client"
	"gandalf/connectors/gojenkins/job"
	"gandalf/libraries/goclient"
	"os"
	"shoset/msg"
	"strconv"

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

		worker.RegisterCommandsFuncs("BUILD_JOB", BuildJob)
		worker.RegisterCommandsFuncs("GET_LAST_SUCCESSFUL_BUILD", GetLastSuccessfulBuild)

		worker.Run()
	}
}

type InputPayload struct {
	Username string
	Password string
	URL      string
	//....
}

func BuildJob(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	fmt.Println("BUILD JOB")
	var buildJobPayload job.BuildJobPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &buildJobPayload)
	if err == nil {
		var clientJenkins *gojenkins.Jenkins
		if buildJobPayload.URL != "" {
			if buildJobPayload.Username != "" && buildJobPayload.Password != "" {
				clientJenkins = client.ClientWithAuthentication(buildJobPayload.URL, buildJobPayload.Username, buildJobPayload.Password)

			} else {
				clientJenkins = client.ClientWithoutAuthentication(buildJobPayload.URL)
			}
		} else {
			clientJenkins = context["client"].(*gojenkins.Jenkins)
		}
		fmt.Println("clientGithub")
		fmt.Println(clientJenkins)

		number, err := job.BuildJob(clientJenkins, buildJobPayload.JobName, buildJobPayload.Params)
		fmt.Println("err")
		fmt.Println(err)

		if err == nil {
			options := make(map[string]string)
			options["payload"] = strconv.Itoa(int(number))
			clientGandalf.SendEvent("JENKINS", "BUILD_JOB", options)
			return 0
		}

	}
	return 1
}

func GetLastSuccessfulBuild(context map[string]interface{}, clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
	fmt.Println("GET LAST SUCCESSFUL BUILD")
	var getLastSuccessfulBuildPayload job.GetLastSuccessfulBuildPayload
	err := json.Unmarshal([]byte(command.GetPayload()), &getLastSuccessfulBuildPayload)
	if err == nil {
		var clientJenkins *gojenkins.Jenkins
		if getLastSuccessfulBuildPayload.URL != "" {
			if getLastSuccessfulBuildPayload.Username != "" && getLastSuccessfulBuildPayload.Password != "" {
				clientJenkins = client.ClientWithAuthentication(getLastSuccessfulBuildPayload.URL, getLastSuccessfulBuildPayload.Username, getLastSuccessfulBuildPayload.Password)

			} else {
				clientJenkins = client.ClientWithoutAuthentication(getLastSuccessfulBuildPayload.URL)
			}
		} else {
			clientJenkins = context["client"].(*gojenkins.Jenkins)
		}
		fmt.Println("clientGithub")
		fmt.Println(clientJenkins)

		_, result, err := job.GetLastSuccessfulBuild(clientJenkins, getLastSuccessfulBuildPayload.JobName)
		fmt.Println("err")
		fmt.Println(err)

		if err == nil {
			options := make(map[string]string)
			options["payload"] = result
			clientGandalf.SendEvent("JENKINS", "GET_LAST_SUCCESSFUL_BUILD", options)
			return 0
		}

	}
	return 1
}
