package cmd

import (
	"database/sql"
	"fmt"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"github.com/liujianping/scaffold/symbol"
	"github.com/revel/revel"
)

func revel_db(project_dir string) (*sql.DB, error) {
	CONFIG, err := revel.LoadConfig("app.conf")
	if err != nil {
		return nil, err
	}

	driver := CONFIG.StringDefault("db.driver", "mysql")
	host := CONFIG.StringDefault("db.host", "127.0.0.1")
	port := CONFIG.IntDefault("db.port", 3306)
	user := CONFIG.StringDefault("db.username", "mym_dev_user")
	pass := CONFIG.StringDefault("db.password", "mym_dev_pass")
	name := CONFIG.StringDefault("db.database", "mym_dev")

	dsn := symbol.DSNFormat(host, port, user, pass, name)
	return sql.Open(driver, dsn)
}

func revel_index(project string, template_dir string, theme string, force bool) error {
	root_dir, err := gopath_src_dir()
	if err != nil {
		return err
	}
	project_dir := path.Join(root_dir, project)
	revel.ConfPaths = append(revel.ConfPaths, path.Join(project_dir, "conf"))

	db, err := revel_db(project_dir)
	if err != nil {
		return err
	}
	defer db.Close()

	tables, err := symbol.GetAllTables(db)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"project": project,
		"tables":  tables,
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "conf", "routes"),
		path.Join(project_dir, "conf", "routes"), data, true); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "views", theme, "header.html"),
		path.Join(project_dir, "app", "views", "header.html"),
		data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "views", theme, "footer.html"),
		path.Join(project_dir, "app", "views", "footer.html"),
		data, force); err != nil {
		return err
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "views", theme, "index.html"),
		path.Join(project_dir, "app", "views", "index.html"),
		data, true); err != nil {
		return err
	}
	return nil
}

func revel_module(project string, template_dir string, moduels []string, theme string, force bool) error {
	if err := revel_model(project, template_dir, moduels, force); err != nil {
		return err
	}

	if err := revel_controller(project, template_dir, moduels, force); err != nil {
		return err
	}

	if err := revel_view(project, template_dir, moduels, theme, force); err != nil {
		return err
	}
	return nil
}

func revel_model(project string, template_dir string, moduels []string, force bool) error {
	root_dir, err := gopath_src_dir()
	if err != nil {
		return err
	}
	project_dir := path.Join(root_dir, project)
	revel.ConfPaths = append(revel.ConfPaths, path.Join(project_dir, "conf"))

	db, err := revel_db(project_dir)
	if err != nil {
		return err
	}
	defer db.Close()

	for _, name := range moduels {
		tables := []*symbol.Table{}
		if name == "*" {
			tbs, err := symbol.GetAllTables(db)
			if err != nil {
				return err
			}
			tables = append(tables, tbs...)
		} else {
			tb, err := symbol.GetTable(db, name)
			if err != nil {
				return err
			}
			tables = append(tables, tb)
		}

		for _, table := range tables {
			data := map[string]interface{}{
				"project": project,
				"table":   table,
			}

			if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "models", "model.crud.go"),
				path.Join(project_dir, "app", "models", fmt.Sprintf("model.%s.go", symbol.ModuleName(table.Name()))),
				data, force); err != nil {
				return err
			}
		}
	}

	return nil
}

func revel_controller(project string, template_dir string, modules []string, force bool) error {
	root_dir, err := gopath_src_dir()
	if err != nil {
		return err
	}
	project_dir := path.Join(root_dir, project)
	revel.ConfPaths = append(revel.ConfPaths, path.Join(project_dir, "conf"))

	db, err := revel_db(project_dir)
	if err != nil {
		return err
	}
	defer db.Close()

	for _, name := range modules {
		tables := []*symbol.Table{}
		if name == "*" {
			tbs, err := symbol.GetAllTables(db)
			if err != nil {
				return err
			}
			tables = append(tables, tbs...)
		} else {
			tb, err := symbol.GetTable(db, name)
			if err != nil {
				return err
			}
			tables = append(tables, tb)
		}

		for _, table := range tables {
			data := map[string]interface{}{
				"project": project,
				"table":   table,
			}

			if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "controllers", "controller.crud.go"),
				path.Join(project_dir, "app", "controllers", fmt.Sprintf("controller.%s.go", symbol.ModuleName(table.Name()))),
				data, force); err != nil {
				return err
			}
		}
	}

	return nil
}

func revel_view(project string, template_dir string, modules []string, theme string, force bool) error {
	root_dir, err := gopath_src_dir()
	if err != nil {
		return err
	}
	project_dir := path.Join(root_dir, project)
	revel.ConfPaths = append(revel.ConfPaths, path.Join(project_dir, "conf"))

	db, err := revel_db(project_dir)
	if err != nil {
		return err
	}
	defer db.Close()

	for _, name := range modules {
		tables := []*symbol.Table{}
		if name == "*" {
			tbs, err := symbol.GetAllTables(db)
			if err != nil {
				return err
			}
			tables = append(tables, tbs...)
		} else {
			tb, err := symbol.GetTable(db, name)
			if err != nil {
				return err
			}
			tables = append(tables, tb)
		}

		for _, table := range tables {
			data := map[string]interface{}{
				"project": project,
				"table":   table,
			}

			if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "views", theme, "crud/create.html"),
				path.Join(project_dir, "app", "views", symbol.ModuleName(table.Name()), "create.html"),
				data, force); err != nil {
				return err
			}

			if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "views", theme, "crud/detail.html"),
				path.Join(project_dir, "app", "views", symbol.ModuleName(table.Name()), "detail.html"),
				data, force); err != nil {
				return err
			}

			if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "views", theme, "crud/index.html"),
				path.Join(project_dir, "app", "views", symbol.ModuleName(table.Name()), "index.html"),
				data, force); err != nil {
				return err
			}

			if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "views", theme, "crud/query.html"),
				path.Join(project_dir, "app", "views", symbol.ModuleName(table.Name()), "query.html"),
				data, force); err != nil {
				return err
			}

			if err := symbol.RenderTemplate(path.Join(template_dir, "revel", "views", theme, "crud/update.html"),
				path.Join(project_dir, "app", "views", symbol.ModuleName(table.Name()), "update.html"),
				data, force); err != nil {
				return err
			}
		}
	}

	return nil
}
