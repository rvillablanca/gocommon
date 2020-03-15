package tmpl

import (
	"bytes"
	"testing"
)

func TestData(t *testing.T) {
	data := NewTD()
	data.Error("an error")
	data.Error("other error")
	data.Info("an info message")
	hasError := data.HasErrors()

	if !hasError {
		t.Error("data should have errors")
	}

	if len(data.Errors) != 2 {
		t.Error("data errors should be 2")
	}

	if len(data.Infos) != 1 {
		t.Error("data info messages should be 1")
	}
}

func TestProvider(t *testing.T) {
	templater := NewTemplater(false)
	templater.AddTemplate("one", "./example.tmpl")
	templater.AddTemplate("two", "./example2.tmpl")

	var w bytes.Buffer
	templater.Render("one", nil, &w)

	content := string(w.Bytes())
	if content != "example" {
		t.Error("invalid content for template one")
	}

	w = bytes.Buffer{}
	templater.Render("two", nil, &w)
	content = string(w.Bytes())
	if content != "example2" {
		t.Error("invalid content for template two")
	}
}
