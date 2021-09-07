package main

import (
	"bufio"
	"fmt"
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
	fmt.Println("Start")

	worker.Start()
	fmt.Println("Start 2")

	clientGandalf := worker.GetClientGandalf()
	fmt.Println("clientGandalf")
	fmt.Println(clientGandalf)
	fmt.Println("Start 3")

	Workflow(clientGandalf)
}
