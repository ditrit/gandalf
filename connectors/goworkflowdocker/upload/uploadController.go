package upload

import (
	"github.com/docker/docker/client"

	"github.com/ditrit/gandalf/connectors/goworkflowdocker/upload/controllers"
)

// Controllers :
type Controllers struct {
	UploadController *controllers.UploadController
}

// ReturnControllers :
func ReturnControllers(cli *client.Client, identity, timeout string, connections []string) *Controllers {

	uploadControllers := new(Controllers)

	uploadControllers.UploadController = controllers.NewUploadController(cli, identity, timeout, connections)

	return uploadControllers
}
