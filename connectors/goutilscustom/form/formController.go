package form

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	goclient "github.com/ditrit/gandalf/libraries/goclient"
	models "github.com/ditrit/gandalf/libraries/goclient/models"

	goformit "github.com/kirves/go-form-it"
)

type FormController struct {
	form          *goformit.Form
	clientGandalf *goclient.ClientGandalf
}

func NewFormController(form *goformit.Form, clientGandalf *goclient.ClientGandalf) *FormController {
	formController := new(FormController)
	formController.form = form
	formController.clientGandalf = clientGandalf

	return formController
}

func (fc FormController) Form(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		fmt.Println("GET")
		tmpl := template.New("name")
		tmpl = template.Must(tmpl.ParseFiles("form/tmpl/layout.tmpl", "form/tmpl/content.tmpl"))
		fmt.Println(fc.form.Render())
		tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"form": fc.form, "Title": "Form"})

	} else if r.Method == "POST" {
		fmt.Println("POST")
		err := r.ParseForm()

		if err != nil {
			tmpl := template.New("name")
			tmpl = template.Must(tmpl.ParseFiles("form/tmpl/layout.tmpl", "form/tmpl/content.tmpl"))
			fmt.Println(fc.form.Render())
			tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"form": fc.form, "Title": "Form"})
		} else {
			for key, value := range r.PostForm {
				fmt.Println(key)
				fmt.Println(value)
			}

			jsonString, err := json.Marshal(r.PostForm)
			if err == nil {
				fc.clientGandalf.SendReply("VALIDATION_FORM", "SUCCES", r.PostForm["UUID"][0], models.NewOptions("", string(jsonString)))
				tmpl := template.New("name")
				tmpl = template.Must(tmpl.ParseFiles("form/tmpl/layout.tmpl", "form/tmpl/succes.tmpl"))
				tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "Validation Succes"})
			} else {
				fc.clientGandalf.SendReply("VALIDATION_FORM", "FAIL", r.PostForm["UUID"][0], models.NewOptions("", ""))
				tmpl := template.New("name")
				tmpl = template.Must(tmpl.ParseFiles("form/tmpl/layout.tmpl", "form/tmpl/fail.tmpl"))
				tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "Validation Fail"})
			}
		}
	}
}

/*
func (fc FormController) FormGet(w http.ResponseWriter, r *http.Request) {

}

func (fc FormController) FormPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	err := r.ParseForm()

	if err != nil {
		tmpl := template.New("name")
		tmpl = template.Must(tmpl.ParseFiles("form/tmpl/layout.tmpl", "form/tmpl/content.tmpl"))
		fmt.Println(fc.form.Render())
		tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"form": fc.form, "Title": "Form"})
	} else {
		jsonString, err := json.Marshal(r.PostForm)
		if err != nil {
			fc.clientGandalf.SendEvent(r.PostForm["uuid"][0], "SUCCES", models.NewOptions("", string(jsonString)))
			tmpl := template.New("name")
			tmpl = template.Must(tmpl.ParseFiles("form/tmpl/layout.tmpl", "form/tmpl/succes.tmpl"))
			tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "Succes"})
		} else {
			fc.clientGandalf.SendEvent(r.PostForm["uuid"][0], "FAIL", models.NewOptions("", ""))
			tmpl := template.New("name")
			tmpl = template.Must(tmpl.ParseFiles("form/tmpl/layout.tmpl", "form/tmpl/fail.tmpl"))
			tmpl.ExecuteTemplate(w, "layout", map[string]interface{}{"Title": "Fail"})
		}
	}

} */
