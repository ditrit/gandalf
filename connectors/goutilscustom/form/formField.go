package form

type Field struct {
	Name     string
	HtmlType string
	Value    interface{}
}

func NewField(name, htmltype string, value interface{}) *Field {
	field := new(Field)
	field.Name = name
	field.HtmlType = htmltype
	field.Value = value

	return field
}

func (f Field) GetName() string {
	return f.Name
}

func (f Field) GetHtmlType() string {
	return f.HtmlType
}

func (f Field) GetValue() interface{} {
	return f.Value
}
