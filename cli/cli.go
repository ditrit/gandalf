package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"test_cli/cli/client"
	cmodels "test_cli/cli/models"
)

func main() {

	var agent string
	flag.StringVar(&agent, "agent", "gandalf", "a string var")

	var typeDB string
	flag.StringVar(&typeDB, "typeDB", "", "a string var")

	var models string
	flag.StringVar(&models, "models", "", "a string var")

	var command string
	flag.StringVar(&command, "command", "", "a string var")

	var value string
	flag.StringVar(&value, "value", "", "a string var")

	flag.Parse()

	fmt.Println("agent:", agent)
	fmt.Println("typeDB:", typeDB)
	fmt.Println("models:", models)
	fmt.Println("command:", command)
	fmt.Println("value:", value)

	test := client.NewClient(agent)

	if typeDB == "gandalf" {
		switch models {
		case "authentication":
			var user cmodels.User
			err := json.Unmarshal([]byte(value), &user)
			fmt.Println(err)
			titi, _ := test.GandalfAuthenticationService.Login(user)
			fmt.Println("TITI")
			fmt.Println(titi)
		case "cluster":
			switch command {
			case "list":
				fmt.Println("lIST")
				var toto []cmodels.Cluster
				toto, _ = test.GandalfClusterService.List(value)
				fmt.Println("PRINT")
				fmt.Println(toto[0].Name)
				fmt.Println(toto[0].Secret)
			}
		}
	} else {

	}

}
