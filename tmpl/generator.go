package tmpl

import (
	"bytes"
	"fmt"
	"html/template"
)

func CreateFromBytes(templates [][]byte) (*template.Template, error) {
	var b bytes.Buffer
	for _, v := range templates {
		b.Write(v)
	}
	var str = b.String()
	t, err := template.New("").Parse(str)
	if err != nil {
		return nil, fmt.Errorf("could not create template: %w", err)
	}

	return t, nil
}
