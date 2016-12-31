package vkclient

import (
	"api"
	"ui"
)

const (
	DialogsHeaders = iota
	Dialog
)

type VkClient struct {
	Ui               *ui.Ui
	Api              *api.Api
	Dialogs          []api.Dialog
	CurrOpenedDialog *api.Dialog
	State            int
}

func NewVkClient(token string) *VkClient {
	client := &VkClient{}

	client.Ui = ui.NewUi()
	if client.Ui == nil {
		return nil
	}

	client.Api = api.NewApi(token)
	if client.Api == nil {
		return nil
	}

	return client
}

func (this *VkClient) Routine() {
	this.setCommandModeHandlers()
	this.Ui.Routine()
}
