package main

import (
	"api"
	"ui"
)

type VkClient struct {
	Ui  *ui.Ui
	Api *api.Api
	// Dialogs []Dialogs
}

func NewVkClient() *VkClient {
	return &VkClient{
		Ui:  ui.NewUi(),
		Api: api.NewApi(),
	}
}

func main() {
	client := NewVkClient()

	client.Api.RequestDialogsHeaders()
	client.Ui.Start()
	client.Ui.CloseUi()
}
