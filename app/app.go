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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	application.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search ips on the internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "search the name of the servers",
			Flags:  flags,
			Action: searchServers,
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

func searchServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	//index, value
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
