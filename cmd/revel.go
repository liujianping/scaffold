package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func revelInit(ctx *cli.Context) {
	template_dir := ctx.GlobalString("template")
	if template_dir == "" {
		fmt.Println("unknown template directory, please use -t to provide.")
		os.Exit(1)
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) == 0 {
		fmt.Println("Usage: scaffold revel init <project>")
		return
	}

	for _, project := range args {
		if err := revel_init(project, template_dir, force); err != nil {
			fmt.Println("scaffold revel init <"+project+"> failed:", err)
			continue
		}
	}
}

func revelIndex(ctx *cli.Context) {
	template_dir := ctx.GlobalString("template")
	if template_dir == "" {
		fmt.Println("unknown template directory, please use -t to provide.")
		return
	}
	force := ctx.GlobalBool("force")
	theme := ctx.String("theme")
	if theme == "" {
		fmt.Println("unknown template theme, please use -t to provide.")
		return
	}

	args := ctx.Args()
	if len(args) == 0 {
		fmt.Println("Usage: scaffold revel index <project>")
		return
	}

	for _, project := range args {
		if err := revel_index(project, template_dir, theme, force); err != nil {
			fmt.Println("scaffold revel index <"+project+"> failed:", err)
			continue
		}
	}
}

func revelModule(ctx *cli.Context) {
	template_dir := ctx.GlobalString("template")
	if template_dir == "" {
		fmt.Println("unknown template directory, please use -t to provide.")
		os.Exit(1)
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) == 0 {
		fmt.Println("Usage: scaffold revel module <project> <module>")
		return
	}

	project := args[0]
	if project == "" {
		fmt.Println("Usage: scaffold revel module <project>")
		return
	}

	var module string = "*"
	if len(args) == 2 {
		module = args[1]
	}

	theme := ctx.String("theme")
	if err := revel_module(project, template_dir, module, theme, force); err != nil {
		fmt.Println("scaffold revel module failed:", err)
		os.Exit(1)
	}
}

func revelController(ctx *cli.Context) {
	template_dir := ctx.GlobalString("template")
	if template_dir == "" {
		fmt.Println("unknown template directory, please use -t to provide.")
		os.Exit(1)
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) == 0 {
		fmt.Println("Usage: scaffold revel controller <project> <module>")
		return
	}

	project := args[0]
	if project == "" {
		fmt.Println("Usage: scaffold revel controller <project> <module>")
		return
	}

	var module string = "*"
	if len(args) == 2 {
		module = args[1]
	}

	if err := revel_controller(project, template_dir, module, force); err != nil {
		fmt.Println("scaffold revel controller failed:", err)
		os.Exit(1)
	}
}

func revelModel(ctx *cli.Context) {
	template_dir := ctx.GlobalString("template")
	if template_dir == "" {
		fmt.Println("unknown template directory, please use -t to provide.")
		os.Exit(1)
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) == 0 {
		fmt.Println("Usage: scaffold revel model <project> <module>")
		return
	}

	project := args[0]
	if project == "" {
		fmt.Println("Usage: scaffold revel model <project> <module>")
		return
	}

	var module string = "*"
	if len(args) == 2 {
		module = args[1]
	}

	if err := revel_model(project, template_dir, module, force); err != nil {
		fmt.Println("scaffold revel model failed:", err)
		os.Exit(1)
	}
}

func revelView(ctx *cli.Context) {
	template_dir := ctx.GlobalString("template")
	if template_dir == "" {
		fmt.Println("unknown template directory, please use -t to provide.")
		os.Exit(1)
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) == 0 {
		fmt.Println("Usage: scaffold revel view <project> <module>")
		return
	}

	project := args[0]
	if project == "" {
		fmt.Println("Usage: scaffold revel view <project> <module>")
		return
	}

	var module string = "*"
	if len(args) == 2 {
		module = args[1]
	}
	theme := ctx.String("theme")

	if err := revel_view(project, template_dir, module, theme, force); err != nil {
		fmt.Println("scaffold revel view failed:", err)
		os.Exit(1)
	}
}
