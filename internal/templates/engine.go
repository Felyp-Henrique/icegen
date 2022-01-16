package templates

import (
	"io"
	"text/template"

	"github.com/Felyp-Henrique/icegen/pkg/templates"
)

type EngineFile struct {
	context  templates.Context
	template *template.Template
}

func NewEngineFile() *EngineFile {
	return &EngineFile{
		template: template.New("simple"),
	}
}

func (e *EngineFile) SetContext(context templates.Context) {
	e.context = context
}
func (e *EngineFile) GetContext() *templates.Context {
	return &e.context
}

func (e *EngineFile) Proccess(writer io.Writer, templateText string) error {
	if temp, err := e.template.Parse(templateText); err == nil {
		return temp.Execute(writer, e.context)
	} else {
		return err
	}
}
