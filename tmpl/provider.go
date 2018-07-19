package tmpl

import (
	"fmt"
	"html/template"
)

//TemplateProvider is responsible to find a template for a given template name.
type TemplateProvider interface {
	FindTemplate(string) *template.Template
	AddTemplate(string, ...string)
}

// DefaultTemplateProvider is the default template provider used
type DefaultTemplateProvider struct {
	templates map[string]*template.Template
}

// FindTemplate finds a template by the given name
func (t DefaultTemplateProvider) FindTemplate(name string) *template.Template {
	tmpl, ok := t.templates[name]
	if !ok {
		panic(fmt.Errorf("could not get template with name %v", name))
	}
	return tmpl
}

// AddTemplate creates a new template with the given name parsing all the files
func (t DefaultTemplateProvider) AddTemplate(name string, files ...string) {
	tmpl := template.Must(template.ParseFiles(files...))
	t.templates[name] = tmpl
}

var _ TemplateProvider = DefaultTemplateProvider{}
