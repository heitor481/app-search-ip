package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	application := cli.NewApp()
	application.Name = "IP Connect"
	application.Usage = "Just a simple terminal app that catches the ip an dns from sites"
	return application
}
