package cmd

import (
	"errors"
	"fmt"
	"go/build"
	"os"
	"path"
	"path/filepath"

	"github.com/codegangsta/cli"
)

func default_template_dir(ctx *cli.Context) (string, error) {
	template_dir := ctx.GlobalString("template")
	if template_dir != "" {
		return template_dir, nil
	}

	gopath := build.Default.GOPATH
	if gopath == "" {
		return "", errors.New("Abort: GOPATH environment variable is not set. ")
	}

	return path.Join(filepath.SplitList(gopath)[0], "src", "github.com/liujianping/scaffold", "template"), nil
}

func gopath_src_dir() (string, error) {
	gopath := build.Default.GOPATH
	if gopath == "" {
		return "", errors.New("Abort: GOPATH environment variable is not set. ")
	}

	return path.Join(filepath.SplitList(gopath)[0], "src"), nil
}

func revelInit(ctx *cli.Context) {
	template_dir, err := default_template_dir(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	force := ctx.GlobalBool("force")
	args := ctx.Args()
	if len(args) == 0 {
		fmt.Println("Usage: scaffold revel init <project>")
		return
	}

	if err := revel_init(args[0], template_dir, force); err != nil {
		fmt.Println("scaffold revel init <"+args[0]+"> failed:", err)
		return
	}
}

func revelIndex(ctx *cli.Context) {
	template_dir, err := default_template_dir(ctx)
	if err != nil {
		fmt.Println(err.Error())
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

	if err := revel_index(args[0], template_dir, theme, force); err != nil {
		fmt.Println("scaffold revel index <"+args[0]+"> failed:", err)
		return
	}

}

func revelModule(ctx *cli.Context) {
	template_dir, err := default_template_dir(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) <= 1 {
		fmt.Println("Usage: scaffold revel module <project> <module> ...")
		return
	}

	project := args[0]
	theme := ctx.String("theme")
	if err := revel_module(project, template_dir, args[1:], theme, force); err != nil {
		fmt.Println("scaffold revel module failed:", err)
		os.Exit(1)
	}
}

func revelController(ctx *cli.Context) {
	template_dir, err := default_template_dir(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) <= 1 {
		fmt.Println("Usage: scaffold revel controller <project> <module> ...")
		return
	}

	project := args[0]
	if err := revel_controller(project, template_dir, args[1:], force); err != nil {
		fmt.Println("scaffold revel controller failed:", err)
		os.Exit(1)
	}
}

func revelModel(ctx *cli.Context) {
	template_dir, err := default_template_dir(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) <= 1 {
		fmt.Println("Usage: scaffold revel model <project> <module> ...")
		return
	}
	project := args[0]

	if err := revel_model(project, template_dir, args[1:], force); err != nil {
		fmt.Println("scaffold revel model failed:", err)
		os.Exit(1)
	}
}

func revelView(ctx *cli.Context) {
	template_dir, err := default_template_dir(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	force := ctx.GlobalBool("force")

	args := ctx.Args()
	if len(args) <= 1 {
		fmt.Println("Usage: scaffold revel view <project> <module> ...")
		return
	}

	project := args[0]
	theme := ctx.String("theme")

	if err := revel_view(project, template_dir, args[1:], theme, force); err != nil {
		fmt.Println("scaffold revel view failed:", err)
		os.Exit(1)
	}
}
