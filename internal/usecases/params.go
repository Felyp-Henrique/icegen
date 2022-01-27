package usecases

import (
	"github.com/Felyp-Henrique/icegen/pkg/templates"
	"github.com/Felyp-Henrique/icegen/pkg/usecases"
)

type CreateParams struct {
	manager templates.Manager
}

func NewCreateParams() *CreateParams {
	return &CreateParams{}
}

func (p *CreateParams) GetManager() templates.Manager {
	return p.manager
}

func (p *CreateParams) SetManager(manager templates.Manager) {
	p.manager = manager
}

func (p *CreateParams) ToParams() usecases.Params {
	return usecases.Params{
		"manager": p.manager,
	}
}
