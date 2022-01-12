package engine

type Engine struct {
	Path string
}

func NewEngine(path string) *Engine {
	return &Engine{
		Path: path,
	}
}

func (e *Engine) Process(template string) string {
	return ""
}
