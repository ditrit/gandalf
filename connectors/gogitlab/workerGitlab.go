package main

import (
	"connectors/gogitlab/workers"
	"encoding/json"
	"os"
)

func main() {
	test()

	/* 	var configuration Configuration
	   	mydir, _ := os.Getwd()
	   	file, _ := os.Open(mydir + "/test.json")
	   	decoder := json.NewDecoder(file)
	   	decoder.Decode(&configuration)

	   	done := make(chan bool)
	   	workerServer := workers.NewWorkerServer(configuration.Identity, configuration.Connections)
	   	go workerServer.Run()
	   	workerProject := workers.NewWorkerProjectBasic(configuration.Identity, configuration.Mail, configuration.Pwd, configuration.Address, configuration.Connections, workerServer)
	   	go workerProject.Run()
	 	//workerhook := workers.NewWorkerHook(configuration.Identity, configuration.Token, configuration.Connections)
	   	//go workerhook.Run()
	<-done */
}

type Configuration struct {
	Identity    string
	Token       string
	Connections []string
	Address     string
	Pwd         string
	Mail        string
}

func test() {

	var configuration Configuration
	mydir, _ := os.Getwd()
	file, _ := os.Open(mydir + "/test.json")
	decoder := json.NewDecoder(file)
	decoder.Decode(&configuration)

	done := make(chan bool)

	workerServer := workers.NewWorkerServer(configuration.Identity, configuration.Connections)
	workerServer.SendEventTest()
	/* 	var configuration Configuration
	   	mydir, err := os.Getwd()
	   	file, _ := os.Open(mydir + "/test.json")

	   	decoder := json.NewDecoder(file)
	   	decoder.Decode(&configuration)

	   	fmt.Println(configuration.Mail)
	   	fmt.Println(configuration.Pwd)
	   	fmt.Println(configuration.Address)
	   	workerProject := workers.NewWorkerProjectTEST(configuration.Mail, configuration.Pwd, configuration.Address, workerServer)

	   	workerProject.AddProjectHookTEST("2", "token", true, false, false, false, false, false, false, false, false, false, false)

	   	if err != nil {
	   		log.Fatal(err)
	   	} */

	//go workerhook.Run()
	<-done
}
