package tmpl

import (
	"io"
)

//Renderer is responsible to write all html templates.
type Renderer interface {
	Render(string, interface{}, io.Writer)
}

// DefaultRenderer is the default Renderer implementation
type DefaultRenderer struct {
	provider TemplateProvider
}

//Render execute the named template on the given writer
func (dr *DefaultRenderer) Render(name string, data interface{}, w io.Writer) {
	template := dr.provider.FindTemplate(name)
	template.Execute(w, data)
}

var (
	_ Renderer = &DefaultRenderer{}
)
