package form

import (
	"net/http"

	"github.com/go-chi/chi"
)

func GetRouter(formController *FormController) *chi.Mux {

	url_patterns := ReturnURLS()

	mux := chi.NewRouter()

	mux.Group(func(mux chi.Router) {
		//FORM
		mux.MethodFunc("GET", url_patterns.FORM_PATH_GET, formController.Form)
		mux.MethodFunc("POST", url_patterns.FORM_PATH_POST, formController.Form)
	})
	return mux
}

func GetRouterWithUrl(formUrl string, formController *FormController) *chi.Mux {
	mux := chi.NewRouter()
	mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./form/tmpl/images/"))))
	mux.Group(func(mux chi.Router) {
		//FORM
		mux.MethodFunc("GET", formUrl, formController.Form)
		mux.MethodFunc("POST", formUrl, formController.Form)
	})
	return mux
}
