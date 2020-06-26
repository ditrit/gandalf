package functions

import (
	"flag"
	"strings"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
)

//Start
func Start() *goclient.ClientGandalf {
	flag.Parse()
	args := flag.Args()
	return goclient.NewClientGandalf(args[0], args[1], strings.Split(args[2], ","))
}
