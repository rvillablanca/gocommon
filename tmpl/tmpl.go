package tmpl

import (
	"io"
)
var Dynamic = false
var provider = NewTemplateProvider(Dynamic)
var renderer = NewRenderer(provider)

func ConfigureDynamic(dynamic bool) {
	Dynamic = dynamic
	provider = NewTemplateProvider(Dynamic)
	renderer = NewRenderer(provider)
}

func SetDelims(left, right string) {
	provider.SetDelims(left, right)
}

// AddTemplate creates a new template with the given name parsing all the files
func AddTemplate(name string, files ...string) error {
	return provider.AddTemplate(name, files...)

}

//Render execute the named template on the given writer
func Render(name string, data interface{}, w io.Writer) error {
	return renderer.Render(name, data, w)
}