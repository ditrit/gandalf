package server

import (
	"github.com/go-chi/chi"
)

func GetRouterWithUrl(serverUrl string, serverController *ServerController) *chi.Mux {
	mux := chi.NewRouter()

	mux.Group(func(mux chi.Router) {
		//FORM
		mux.MethodFunc("POST", serverUrl, serverController.Hook)
	})
	return mux
}
