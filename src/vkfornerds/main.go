package main

import (
	"api"
	"flag"
	"log"
	"ui"
)

type VkClient struct {
	Ui  *ui.Ui
	Api *api.Api
	// Dialogs []Dialog
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

	tmp := client.Api.RequestDialogsHeaders()
	log.Println(tmp)
	client.Ui.Start()
	client.Ui.CloseUi()
}
