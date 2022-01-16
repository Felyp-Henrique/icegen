package usecases

import (
	"github.com/Felyp-Henrique/icegen/pkg/templates"
	"github.com/Felyp-Henrique/icegen/pkg/usecases"
)

func UseCaseCreate(params usecases.Params) {
	var manager templates.Manager = params["manager"].(templates.Manager)

	if err := manager.Proccess("icecast.xml"); err != nil {
		panic(err)
	}
}
