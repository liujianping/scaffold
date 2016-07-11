package commands

import (
	"github.com/codegangsta/cli"
)

const version = "1.0.1"
const author = ""
const support = "liujianping.itech@qq.com"

func App() *cli.App {
	app := cli.NewApp()

	//! app settings
	app.Name = "scaffold"
	app.Usage = "scaffolding, generate code by database schema definitions"
	app.Version = version
	app.Author = author
	app.Email = support

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "force to replace file by generated",
		},
		cli.StringFlag{
			Name:  "template, t",
			Value: "portal",
			Usage: "template directory(default: portal)",
		},
		cli.StringFlag{
			Name:  "template-folder",
			Value: "",
			Usage: "custom template abosolute folder path",
		},
		cli.StringSliceFlag{
			Name:  "include-template-suffix, i",
			Value: &cli.StringSlice{},
			Usage: "include template suffix or template name",
		},
		cli.StringSliceFlag{
			Name:  "exclude-template-suffix, x",
			Value: &cli.StringSlice{},
			Usage: "exclude template suffix or template name",
		},
	}
	//! app commands
	app.Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: "generate code from templates",
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
					Usage: "database host(default: localhost)",
				},
				cli.IntFlag{
					Name:  "port, P",
					Value: 3306,
					Usage: "database port(default: 3306)",
				},
				cli.StringFlag{
					Name:  "username, u",
					Value: "root",
					Usage: "database user name(default: root)",
				},
				cli.StringFlag{
					Name:  "password, p",
					Value: "",
					Usage: "database user password",
				},
			},
			Action: func(ctx *cli.Context) {
				Generate(ctx)
			},
		},
	}

	return app
}
