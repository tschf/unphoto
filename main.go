package main

import (
	"log"
	"os"

	"github.com/tschf/unphoto/source/guardian"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "unphoto"
	app.Usage = "Download the latest pic of the day"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "guardian",
			Usage: "Download the pic of the day from The Guardian",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Bool("guardian") {
			guardian.GetPhoto()

		}
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
