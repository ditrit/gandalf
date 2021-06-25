package upload

import (
	"net/http"

	"github.com/docker/docker/client"

	"github.com/gorilla/mux"
)

func GetRouter(cli *client.Client, identity, timeout string, connections []string) *mux.Router {

	//CONTROLLERS
	controllers := ReturnControllers(cli, identity, timeout, connections)

	//URLS
	urls := ReturnURLS()

	mux := mux.NewRouter()
	mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))

	//TODO REVOIR
	mux.HandleFunc(urls.UPLOAD_PATH_GET, controllers.UploadController.Get).Methods("GET")
	mux.HandleFunc(urls.UPLOAD_PATH_POST, controllers.UploadController.Post).Methods("POST")

	return mux
}
