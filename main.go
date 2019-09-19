package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tschf/unphoto/config"
	"github.com/tschf/unphoto/source/guardian"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = config.APP_NAME
	app.Usage = "Download the latest photo of the day"
	app.Version = config.APP_VERSION

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "guardian",
			Usage: "Download the photo of the day from The Guardian",
		},
		cli.BoolFlag{
			Name:  "wallpaper",
			Usage: "Set the pic as the wallpaper",
		},
	}
	app.Action = func(c *cli.Context) error {

		applyWallpaper := c.Bool("wallpaper")

		if c.Bool("guardian") {
			fmt.Println("Photo source: Guardian")
			guardian.GetPhoto(applyWallpaper)

		}
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
