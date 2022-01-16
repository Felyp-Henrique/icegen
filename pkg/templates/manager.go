package templates

type Manager interface {
	SetEngine(engine Engine)
	GetEngine() Engine
	SetPath(path string)
	GetPath() string
	Proccess(fileName string) error
}
