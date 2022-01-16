package usecases

import (
	"github.com/Felyp-Henrique/icegen/pkg/templates"
	"github.com/Felyp-Henrique/icegen/pkg/usecases"
)

type CreateParams struct {
	usecases.Params
}

func NewCreateParams() *CreateParams {
	return &CreateParams{
		Params: usecases.Params{},
	}
}

func (p *CreateParams) GetManager() templates.Manager {
	return p.Params["manager"].(templates.Manager)
}

func (p *CreateParams) SetManager(manager templates.Manager) {
	p.Params["manager"] = manager
}
