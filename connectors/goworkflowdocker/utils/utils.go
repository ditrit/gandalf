package utils

import (
	"fmt"
	"strings"
)

func GetPathFromConnections(slice []string) string {
	fmt.Println("slice")
	fmt.Println(slice)
	splitslice := strings.Split(slice[0], "/")
	fmt.Println("splitslice")
	fmt.Println(splitslice)
	if len(splitslice) > 0 {
		splitslice = splitslice[:len(splitslice)-1]
		fmt.Println("splitslice2")
		fmt.Println(splitslice)
	}
	slicestring := strings.Join(splitslice, "/")
	fmt.Println("slicestring")
	fmt.Println(slicestring)
	return slicestring + "/"
}

func ChangePathFromConnections(slice []string) string {
	splitslice := strings.Split(slice[0], "/")

	return "/var/run/sockets/" + splitslice[len(splitslice)-1]
}

func TransformConnectionsSliceToString(slice []string) string {
	return strings.Join(slice, ",")
}
