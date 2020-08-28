package form

import (
	"fmt"

	goform "github.com/kirves/go-form-it"
	"github.com/kirves/go-form-it/fields"
)

func CreateForm(uuid string, fieldArray []Field) *goform.Form {
	form := goform.BootstrapForm(goform.POST, "/gandalf/form/")
	form.Elements(goform.FieldSet("uuid", fields.HiddenField("UUID").SetValue(uuid)))
	for _, field := range fieldArray {
		switch field.GetHtmlType() {
		case "TextField":
			//form.Elements(fields.TextField(field.GetName()).SetLabel(field.GetName()).AddClass("form-control"))
			form.Elements(fields.TextField(field.GetName()).SetLabel(field.GetName()))
			break
		case "Checkbox":
			form.Elements(fields.Checkbox(field.GetName(), field.GetValue().(bool)))
			break
		case "DatetimeField":
			form.Elements(fields.DatetimeField(field.GetName()))
			break
		case "NumberField":
			form.Elements(fields.NumberField(field.GetName()))
			break
		case "StructField": //TODO REVOIR
			break
		default:
			break
		}
	}
	form.Elements(fields.SubmitButton("btn1", "Submit").AddClass("btn btn-primary"))
	fmt.Println(form)
	return form
}

func CreateFormWithUrl(url, uuid string, fieldArray []Field) *goform.Form {
	fmt.Println("URL")
	fmt.Println(url)
	fmt.Println("TATA")
	form := goform.BootstrapForm(goform.POST, url)
	//form := goform.BaseForm(goform.POST, url)
	fmt.Println("form")
	fmt.Println(form)
	form.Elements(fields.HiddenField("UUID").SetValue(uuid))

	for _, field := range fieldArray {
		switch field.GetHtmlType() {
		case "TextField":
			//form.Elements(fields.TextField(field.GetName()).SetLabel(field.GetName()).AddClass("form-control"))
			form.Elements(fields.TextField(field.GetName()).SetId(field.GetName()).SetLabel(field.GetName() + " : ").AddLabelClass("font-weight-bold").AddLabelClass("text-capitalize"))
			break
		case "Checkbox":
			form.Elements(fields.Checkbox(field.GetName(), field.GetValue().(bool)))
			break
		case "DatetimeField":
			form.Elements(fields.DatetimeField(field.GetName()))
			break
		case "NumberField":
			form.Elements(fields.NumberField(field.GetName()))
			break
		case "StructField": //TODO REVOIR
			break
		default:
			break
		}
	}
	form.Elements(fields.SubmitButton("btn1", "Submit").AddClass("btn btn-primary"))

	return form
}
