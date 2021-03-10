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

	worker.RegisterCommandsFuncs("CREATE_FORM", CreateForm)
	worker.RegisterCommandsFuncs("SEND_AUTH_MAIL", SendAuthMail)

	worker.Run()
}

//TODO REVOIR
func SendAuthMail(clientGandalf *goclient.ClientGandalf, major int64, command msg.Command) int {
}
