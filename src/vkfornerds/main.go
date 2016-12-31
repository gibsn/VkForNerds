package main

import (
	"flag"
	"log"
	"os"

	"vkclient"
)

func RedirectLogsToFile(file string) {
	if file != "" {
		file, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Println(err)
			log.Fatalf("Could not open file %s for logging\n", file)
		}
		log.SetOutput(file)
	}
}

func main() {
	token := flag.String("t", "", "vk token")
	logFile := flag.String("o", "", "log to file")
	flag.Parse()

	RedirectLogsToFile(*logFile)

	if *token == "" {
		log.Fatal("You have not provided token")
	}

	client := vkclient.NewVkClient()
	client.Api.AccessToken = *token

	client.Routine()
}
