package cmd

import (
	"database/sql"
	"fmt"
	"path"

	"github.com/codegangsta/cli"
	_ "github.com/go-sql-driver/mysql"
	"github.com/liujianping/scaffold/symbol"
)

func modelGenerate(ctx *cli.Context) {
	template_dir, err := default_template_dir(ctx)
	if err != nil {
		fmt.Println("scaffold model generate failed:", err)
		return
	}

	args := ctx.Args()
	if len(args) != 1 {
		fmt.Println("Usage: scaffold model generate <project/path>")
		return
	}

	project := args[0]
	root_dir, err := gopath_src_dir()
	if err != nil {
		fmt.Println("scaffold model generate failed:", err)
		return
	}
	project_dir := path.Join(root_dir, project)

	driver := ctx.String("driver")
	database := ctx.String("database")
	host := ctx.String("host")
	port := ctx.Int("port")
	user := ctx.String("username")
	pass := ctx.String("password")

	dsn := symbol.DSNFormat(host, port, user, pass, database)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		fmt.Println("scaffold model generate failed:", err)
		return
	}
	defer db.Close()

	tables, err := symbol.GetAllTables(db)
	if err != nil {
		fmt.Println("scaffold model generate failed:", err)
		return
	}

	if err := symbol.RenderTemplate(path.Join(template_dir, "model", "model.go.t"),
		path.Join(project_dir, "model.go"), map[string]interface{}{
			"project": project,
			"tables":  tables,
		}, true); err != nil {
		fmt.Println("scaffold model generate failed:", err)
		return
	}

	for _, table := range tables {
		data := map[string]interface{}{
			"table": table,
		}

		dest := fmt.Sprintf("model.%s.go", symbol.ModuleName(table.Name()))
		if err := symbol.RenderTemplate(path.Join(template_dir, "model", "model.table.go.t"),
			path.Join(project_dir, dest), data, true); err != nil {
			fmt.Println("scaffold model generate failed:", err)
			return
		}
	}
}
