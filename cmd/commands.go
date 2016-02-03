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
	app.Usage = "scaffold, generate revel project by database schema"
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
			Name:  "model",
			Usage: "data model scaffolding",
			Subcommands: []cli.Command{
				{
					Name:  "generate",
					Usage: "generate model code",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "driver, D",
							Value: "mysql",
							Usage: "database driver, now only support mysql",
						},
						cli.StringFlag{
							Name:  "database, d",
							Value: "",
							Usage: "database name",
						},
						cli.StringFlag{
							Name:  "host, H",
							Value: "localhost",
							Usage: "database host",
						},
						cli.IntFlag{
							Name:  "port, P",
							Value: 3306,
							Usage: "database port",
						},
						cli.StringFlag{
							Name:  "username, u",
							Value: "root",
							Usage: "database user name",
						},
						cli.StringFlag{
							Name:  "password, p",
							Value: "",
							Usage: "database user password",
						},
					},
					Action: func(ctx *cli.Context) {
						modelGenerate(ctx)
					},
				},
			},
		},
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
