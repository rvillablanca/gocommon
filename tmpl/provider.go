package tmpl

import (
	"fmt"
	"html/template"
)

//TemplateProvider is responsible to find a template for a given template name.
type TemplateProvider interface {
	FindTemplate(string) (*template.Template, error)
	AddTemplate(string, ...string) error
}

// DefaultTemplateProvider is the default template provider used
type DefaultTemplateProvider struct {
	templates map[string]*template.Template
}

// FindTemplate finds a template by the given name
func (t *DefaultTemplateProvider) FindTemplate(name string) (*template.Template, error) {
	tmpl, ok := t.templates[name]
	if !ok {
		return nil, fmt.Errorf("could not get template with name %v", name)
	}
	return tmpl, nil
}

// AddTemplate creates a new template with the given name parsing all the files
func (t *DefaultTemplateProvider) AddTemplate(name string, files ...string) error {
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	t.templates[name] = tmpl
	return nil
}

var _ TemplateProvider = &DefaultTemplateProvider{}

// NewTemplateProvider create a new provider
func NewTemplateProvider() TemplateProvider {
	p := new(DefaultTemplateProvider)
	p.templates = make(map[string]*template.Template)
	return p
}
