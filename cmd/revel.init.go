package cmd

import (
	"errors"
	"go/build"
	"path"
	"path/filepath"

	"github.com/liujianping/scaffold/symbol"
)

func revel_init(project string, template_dir string, force bool) error {

	gopath := build.Default.GOPATH
	if gopath == "" {
		return errors.New("Abort: GOPATH environment variable is not set. ")
	}

	// set go src path
	srcRoot := filepath.Join(filepath.SplitList(gopath)[0], "src")
	project_dir := path.Join(srcRoot, project)

	data := map[string]interface{}{
		"project": project,
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "models", "model.go"),
		path.Join(project_dir, "app", "models", "model.go"), data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "models", "model.sql.go"),
		path.Join(project_dir, "app", "models", "model.sql.go"), data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "models", "model.validator.go"),
		path.Join(project_dir, "app", "models", "model.validator.go"), data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "controllers", "controller.db.go"),
		path.Join(project_dir, "app", "controllers", "controller.db.go"), data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "controllers", "controller.go"),
		path.Join(project_dir, "app", "controllers", "controller.go"), data, force); err != nil {
		return err
	}

	return nil
}
