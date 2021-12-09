package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func GetRouter(workflowController *WorkflowController) *chi.Mux {

	url_patterns := ReturnURLS()

	mux := chi.NewRouter()
	mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./tmpl/images/"))))
	mux.Group(func(mux chi.Router) {
		//FORM
		mux.MethodFunc("GET", url_patterns.SEND_COMMAND_PATH, workflowController.SendCommand)
		mux.MethodFunc("GET", url_patterns.SEND_UPDATE_PATH, workflowController.SendUpdate)
		mux.MethodFunc("GET", url_patterns.INDEX_PATH, workflowController.Index)
		mux.MethodFunc("GET", url_patterns.EVENTS_PATH, workflowController.b.ServeHTTP)
	})
	return mux
}
