package usecases

type Params map[string]interface{}

type ToParams interface {
	ToParams() Params
}
