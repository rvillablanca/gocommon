package tmpl

import (
	"io"
)

//Renderer is responsible to write all html templates.
type Renderer interface {
	Render(string, interface{}, io.Writer) error
}

// DefaultRenderer is the default Renderer implementation
type DefaultRenderer struct {
	provider TemplateProvider
}

//Render execute the named template on the given writer
func (dr *DefaultRenderer) Render(name string, data interface{}, w io.Writer) error {
	template, err := dr.provider.FindTemplate(name)
	if err != nil {
		return err
	}
	template.Execute(w, data)
	return nil
}

var (
	_ Renderer = &DefaultRenderer{}
)

// NewRenderer creates a new Renderer
func NewRenderer(provider TemplateProvider) Renderer {
	r := new(DefaultRenderer)
	r.provider = provider
	return r
}
