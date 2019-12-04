package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("Usage : ")
	fmt.Println("	tcp --client host:port")
	fmt.Println("	tcp --server port")
	return
}

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
		clientTcp(connect)
	case "-c":
		clientTcp(connect)
	case "--server":
		serverTcp(connect)
	case "-s":
		serverTcp(connect)
	default:
		usage()
	}

}
