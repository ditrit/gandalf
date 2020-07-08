package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
	models "github.com/ditrit/gandalf/libraries/goclient/models"
)

type AppController struct {
	url           string
	clientGandalf *goclient.ClientGandalf
}

func NewAppController(url string, clientGandalf *goclient.ClientGandalf) *AppController {
	appController := new(AppController)
	appController.url = url
	appController.clientGandalf = clientGandalf

	return appController
}

func (ac AppController) App(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		fmt.Println("GET")
		tmpl := template.New("name")
		tmpl = template.Must(tmpl.ParseFiles("app/tmpl/layout.tmpl", "app/tmpl/content.tmpl"))
		fmt.Println("ac.url")
		fmt.Println(ac.url)
		tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "App", "Url": ac.url})

	} else if r.Method == "POST" {
		fmt.Println("POST")
		err := r.ParseForm()

		if err != nil {
			tmpl := template.New("name")
			tmpl = template.Must(tmpl.ParseFiles("app/tmpl/layout.tmpl", "app/tmpl/content.tmpl"))
			tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "App", "Url": ac.url})
		} else {
			jsonString, err := json.Marshal(r.PostForm["appName"][0])
			if err == nil {
				ac.clientGandalf.SendEvent("Application", "NEW_APP", models.NewOptions("", string(jsonString)))
				tmpl := template.New("name")
				tmpl = template.Must(tmpl.ParseFiles("form/tmpl/layout.tmpl", "form/tmpl/succes.tmpl"))
				tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "Application Succes"})
			}
		}
	}
}
