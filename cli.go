package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nexus-client-go"
	app.Usage = "Client for the Nexus Configuration Server"
	app.Action = func(c *cli.Context) {
		println("See nexus-client-go help for more information")
	}
	app.Author = "Eduardo Trujillo <ed@chromabits.com>"

	app.Run(os.Args)
}
