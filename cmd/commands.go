package cmd

import (
	"github.com/codegangsta/cli"
)

const version = "0.0.1"
const author = ""
const support = "liujianping@h2object.io"

func App() *cli.App {
	app := cli.NewApp()

	//! app settings
	app.Name = "scaffold"
	app.Usage = "scaffold database adminstrator portal tool"
	app.Version = version
	app.Author = author
	app.Email = support

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "template, t",
			Value: "",
			Usage: "template directory",
		},
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "force execute",
		},
	}
	//! app commands
	app.Commands = []cli.Command{
		{
			Name:  "revel",
			Usage: "revel project scaffolding",
			Subcommands: []cli.Command{
				{
					Name:  "init",
					Usage: "revel project scaffolding initialization",
					Action: func(ctx *cli.Context) {
						revelInit(ctx)
					},
				},
				{
					Name:  "index",
					Usage: "revel project scaffolding indexing",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "theme, t",
							Value: "bootstrap",
							Usage: "revel project view theme",
						},
					},
					Action: func(ctx *cli.Context) {
						revelIndex(ctx)
					},
				},
				{
					Name:  "module",
					Usage: "revel project scaffolding module",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "theme, t",
							Value: "bootstrap",
							Usage: "revel project module view theme",
						},
					},
					Action: func(ctx *cli.Context) {
						revelModule(ctx)
					},
				},
				{
					Name:  "controller",
					Usage: "revel project scaffolding controller",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "module, m",
							Value: "*",
							Usage: "revel project module name",
						},
					},
					Action: func(ctx *cli.Context) {
						revelController(ctx)
					},
				},
				{
					Name:  "model",
					Usage: "revel project scaffolding model",
					Action: func(ctx *cli.Context) {
						revelModel(ctx)
					},
				},
				{
					Name:  "view",
					Usage: "revel project scaffolding view",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "theme, t",
							Value: "bootstrap",
							Usage: "revel project module view theme",
						},
					},
					Action: func(ctx *cli.Context) {
						revelView(ctx)
					},
				},
			},
		},
	}

	return app
}
