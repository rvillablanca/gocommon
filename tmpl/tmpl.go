package tmpl

import "io"

var provider DefaultTemplateProvider
var renderer DefaultRenderer

func init() {
	renderer.provider = provider
}

// AddTemplate creates a new template with the given name parsing all the files
func AddTemplate(name string, files ...string) {
	provider.AddTemplate(name, files...)
}

//Render execute the named template on the given writer
func Render(name string, data interface{}, w io.Writer) {
	renderer.Render(name, data, w)
}
