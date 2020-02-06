package main

import (
	"fmt"
	"os"
	"strconv"
)

//usage :
func usage() {
	fmt.Println("Usage : ")
	fmt.Println("	tcp --client host:port")
	fmt.Println("	tcp --server port")
}

//main :
func main() {
	arguments := os.Args
	fmt.Println("nb args:  " + strconv.Itoa(len(arguments)) + ".")

	if len(arguments) != 3 {
		usage()
		return
	}

	option := arguments[1]
	connect := arguments[2]

	switch option {
	case "--client":
	case "-c":
		clientTCP(connect)
	case "--server":
	case "-s":
		serverTCP(connect)
	default:
		usage()
	}
}
