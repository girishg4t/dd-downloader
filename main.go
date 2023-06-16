package main

import (
	"log"

	"github.com/dd-downloader/cmd"
)

func main() {
	cmd.Execute()
	log.Println("Done Downloading")
}
