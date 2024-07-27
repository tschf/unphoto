package main

import (
	"log"
	"os"

	"github.com/tschf/unphoto/config"
	"github.com/tschf/unphoto/source/guardian"
	"github.com/tschf/unphoto/source/local"

	"github.com/urfave/cli"
)

type photoSource interface {
	GetPhoto(applyWallpaper bool)
	PrintSourceInfo()
}

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
			Name:  "local",
			Usage: "Choose a photo from the specified localPath",
		},
		cli.StringFlag{
			Name:  "local-path",
			Usage: "Specify the path to choose a picture. One will be chosen at random",
		},
		cli.BoolFlag{
			Name:  "wallpaper",
			Usage: "Set the pic as the wallpaper",
		},
	}
	app.Action = func(c *cli.Context) error {

		var ps photoSource
		var applyWallpaper bool

		if c.Bool("guardian") {
			ps = guardian.GuardianSource{}
		} else if c.Bool("local") {
			ps = local.LocalSource{}
			ps.(local.LocalSource).SetPicturePath(c.String("local-path"))
		}

		applyWallpaper = c.Bool("wallpaper")
		ps.PrintSourceInfo()
		ps.GetPhoto(applyWallpaper)
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
