package tmpl

import (
	"html/template"
	"io"
)

type Templater interface {
	TemplateProvider
	Renderer
}

type defaultTemplater struct {
	templateProvider TemplateProvider
	renderer         Renderer
}

func (dt *defaultTemplater) FindTemplate(t string) (*template.Template, error) {
	return dt.templateProvider.FindTemplate(t)
}

func (dt *defaultTemplater) AddTemplate(t string, f ...string) error {
	return dt.templateProvider.AddTemplate(t, f...)
}

func (dt *defaultTemplater) SetDelims(l string, r string) {
	dt.templateProvider.SetDelims(l, r)
}

func (dt *defaultTemplater) Render(t string, d interface{}, w io.Writer) error {
	return dt.renderer.Render(t, d, w)
}

func NewTemplater(dynamic bool) Templater {
	var t defaultTemplater
	t.templateProvider = NewTemplateProvider(dynamic)
	t.renderer = NewRenderer(t.templateProvider)
	return &t
}
