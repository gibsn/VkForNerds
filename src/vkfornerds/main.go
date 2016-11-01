package main

import (
	"api"
	"flag"
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
	token := flag.String("t", "", "vk token")
	flag.Parse()

	client := NewVkClient()
	client.Api.AccessToken = *token

	// client.Api.RequestDialogsHeaders()
	client.Ui.Start()
	client.Ui.CloseUi()
}
