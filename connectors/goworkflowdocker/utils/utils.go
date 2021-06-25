package utils

import "strings"

func GetPathFromConnections(slice []string) string {
	splitslice := strings.Split(slice[0], "/")
	if len(splitslice) > 0 {
		splitslice = splitslice[:len(splitslice)-1]
	}
	slicestring := strings.Join(splitslice, "/")
	return slicestring + "/"
}

func TransformConnectionsSliceToString(slice []string) string {
	return strings.Join(slice, ",")
}
