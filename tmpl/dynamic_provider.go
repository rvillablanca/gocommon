package tmpl

import (
	"fmt"
	"html/template"
)

type DynamicProvider struct {
	DefaultTemplateProvider
	dynamics map[string][]string
}

func (dp DynamicProvider) FindTemplate(name string) (*template.Template, error) {
	if _, ok := dp.dynamics[name]; !ok {
		return nil, fmt.Errorf("template %s does not exists", name)
	}

	tmpl, err := template.ParseFiles(dp.dynamics[name]...)
	if err != nil {
		return nil, err
	}

	tmpl.Delims(dp.left, dp.right)
	return tmpl, nil
}

func (dp DynamicProvider) AddTemplate(name string, files...string) error {
	if _, ok := dp.dynamics[name]; ok {
		return fmt.Errorf("template %s already exists", name)
	}

	dp.dynamics[name] = files
	return nil
}

func (dp DynamicProvider) SetDelims(left, right string) {
	dp.DefaultTemplateProvider.SetDelims(left, right)
}

func NewDynamic() DynamicProvider {
	defaultTemplateProvider := NewTemplateProvider(false).(DefaultTemplateProvider)
	provider := DynamicProvider{DefaultTemplateProvider: defaultTemplateProvider}
	provider.dynamics = make(map[string][]string)
	return provider
}




