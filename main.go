package main

import (
	"application/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Panic("Panic")
	}
}
