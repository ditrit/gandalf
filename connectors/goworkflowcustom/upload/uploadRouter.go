package upload

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func GetRouterWithUrl(serverUrl string, uploadController *UploadController) *chi.Mux {
	mux := chi.NewRouter()
	mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./upload/tmpl/images/"))))
	fmt.Println("serverUrl")
	fmt.Println(serverUrl)
	mux.Group(func(mux chi.Router) {
		//FORM
		mux.MethodFunc("GET", serverUrl, uploadController.Upload)
		mux.MethodFunc("POST", serverUrl, uploadController.Upload)
	})
	return mux
}
