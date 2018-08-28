package tmpl

import "io"

var provider = NewTemplateProvider()
var renderer = NewRenderer(provider)

// AddTemplate creates a new template with the given name parsing all the files
func AddTemplate(name string, files ...string) error {
	return provider.AddTemplate(name, files...)

}

//Render execute the named template on the given writer
func Render(name string, data interface{}, w io.Writer) error {
	return renderer.Render(name, data, w)
}
