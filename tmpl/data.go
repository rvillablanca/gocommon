package tmpl

// TemplateData is used as a common data for templates.
type TemplateData struct {
	Data   map[string]interface{}
	Errors []string
	Infos  []string
	URL    string
}

// NewTD allow create an empty TemplateData
func NewTD() *TemplateData {
	td := new(TemplateData)
	td.Errors = make([]string, 0)
	td.Infos = make([]string, 0)
	td.Data = make(map[string]interface{})
	return td
}

// Add allows add an error message to TemplateData
func (td *TemplateData) Add(error string) {
	td.Errors = append(td.Errors, error)
}

// Info allows add an info message to TemplateData
func (td *TemplateData) Info(error string) {
	td.Infos = append(td.Infos, error)
}

// HasErrors allows check if the TemplateData has errors
func (td *TemplateData) HasErrors() bool {
	return len(td.Errors) > 0
}
