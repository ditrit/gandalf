package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func GetRouterWithUrl(serverUrl string, appController *AppController) *chi.Mux {
	mux := chi.NewRouter()
	mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	fmt.Println("serverUrl")
	fmt.Println(serverUrl)
	mux.Group(func(mux chi.Router) {
		//FORM
		mux.MethodFunc("GET", serverUrl, appController.App)
		mux.MethodFunc("POST", serverUrl, appController.App)
	})
	return mux
}
