package upload

import (
	"fmt"
	"net/http"

	"github.com/docker/docker/client"

	"github.com/gorilla/mux"
)

func GetRouter(cli *client.Client, identity, timeout string, connections []string) *mux.Router {
	fmt.Println("toto2")

	//CONTROLLERS
	controllers := ReturnControllers(cli, identity, timeout, connections)
	fmt.Println("toto3")

	//URLS
	urls := ReturnURLS()
	fmt.Println("toto4")

	mux := mux.NewRouter()
	mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	fmt.Println("toto5")

	//TODO REVOIR
	mux.HandleFunc(urls.UPLOAD_PATH_GET, controllers.UploadController.Get).Methods("GET")
	mux.HandleFunc(urls.UPLOAD_PATH_POST, controllers.UploadController.Post).Methods("POST")

	return mux
}
