package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	application := cli.NewApp()
	application.Name = "IP Connect"
	application.Usage = "Just a simple terminal app that catches the ip an dns from sites"

	application.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search ips on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIps,
		},
	}
	return application
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	//function net
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	//index, value
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
