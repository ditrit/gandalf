package server

import (
	"io/ioutil"
	"libraries/goclient"
	"libraries/goclient/models"
	"net/http"
)

type ServerController struct {
	clientGandalf *goclient.ClientGandalf
}

func NewServerController(clientGandalf *goclient.ClientGandalf) *ServerController {
	serverController := new(ServerController)
	serverController.clientGandalf = clientGandalf

	return serverController
}

func (sc ServerController) Hook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		bodyString := string(body)
		sc.clientGandalf.SendEvent("Gitlab", "HOOK", models.NewOptions("", bodyString))
	}
}
