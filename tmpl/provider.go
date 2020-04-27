package tmpl

import (
	"fmt"
	"html/template"
)

//TemplateProvider is responsible to find a template for a given template name.
type TemplateProvider interface {
	FindTemplate(string) (*template.Template, error)
	AddTemplate(string, ...string) error
	SetDelims(string, string)
}

// DefaultTemplateProvider is the default template provider used
type DefaultTemplateProvider struct {
	templates map[string]*template.Template
	left      string
	right     string
}

func (t DefaultTemplateProvider) SetDelims(left, right string) {
	t.left = left
	t.right = right
}

// FindTemplate finds a template by the given name
func (t DefaultTemplateProvider) FindTemplate(name string) (*template.Template, error) {
	tmpl, ok := t.templates[name]
	if !ok {
		return nil, fmt.Errorf("could not get template with name %v", name)
	}
	return tmpl, nil
}

// AddTemplate creates a new template with the given name parsing all the files
func (t DefaultTemplateProvider) AddTemplate(name string, files ...string) error {
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	tmpl.Delims(t.left, t.right)
	t.templates[name] = tmpl
	return nil
}

var _ TemplateProvider = DefaultTemplateProvider{}
var _ TemplateProvider = DynamicProvider{}

// NewTemplateProvider create a new provider
func NewTemplateProvider(dynamic bool) TemplateProvider {
	if dynamic {
		return NewDynamic()
	}

	p := DefaultTemplateProvider{}
	p.templates = make(map[string]*template.Template)
	return p
}