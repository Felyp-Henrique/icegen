package templates

import (
	"io"
	"os"
	"path/filepath"

	"github.com/Felyp-Henrique/icegen/pkg/templates"
)

type ManagerFile struct {
	engine templates.Engine
	path   string
	writer io.Writer
}

func NewManagerFile(engine *EngineFile, path string, output string) (*ManagerFile, error) {
	if file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return nil, err
	} else {
		return &ManagerFile{
			engine: engine,
			path:   path,
			writer: file,
		}, nil
	}
}

func (m *ManagerFile) SetEngine(engine templates.Engine) {
	m.engine = engine
}

func (m *ManagerFile) GetEngine() templates.Engine {
	return m.engine
}

func (m *ManagerFile) SetPath(path string) {
	m.path = path
}

func (m *ManagerFile) GetPath() string {
	return m.path
}

func (m *ManagerFile) Proccess(fileName string) error {
	templatePath := filepath.Join(m.path, fileName)
	if templateBytes, err := os.ReadFile(templatePath); err != nil {
		return err
	} else {
		return m.engine.Proccess(m.writer, string(templateBytes))
	}
}
