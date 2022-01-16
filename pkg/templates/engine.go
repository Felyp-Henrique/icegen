package templates

import "io"

type Engine interface {
	SetContext(context Context)
	GetContext() *Context
	Proccess(writer io.Writer, templateText string) error
}
