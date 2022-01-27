package templates

type Context map[string]interface{}

type ToContext interface {
	ToContext() Context
}
