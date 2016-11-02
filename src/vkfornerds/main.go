package main

import (
	"flag"
	"vkclient"
)

func main() {
	token := flag.String("t", "", "vk token")
	flag.Parse()

	client := vkclient.NewVkClient()
	client.Api.AccessToken = *token

	client.Start()
}
