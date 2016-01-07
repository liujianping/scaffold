package controllers

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"

	"github.com/revel/revel"
	gorp "gopkg.in/gorp.v1"
)

type DBController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *DBController) Index() revel.Result {
	return c.RenderTemplate("index.html")
}

func (c *DBController) Begin() revel.Result {
	txn, err := model.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *DBController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *DBController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *DBController) Pagination(fn string, total, page, size int) template.HTML {
	if total <= size {
		return template.HTML("")
	}
	pages := total / size
	if total%size > 0 {
		pages = pages + 1
	}

	var buffer bytes.Buffer
	buffer.WriteString("<nav>")
	buffer.WriteString("<ul class=\"pagination pagination-sm\">")
	for i := 0; i < pages; i++ {
		if i == page {
			buffer.WriteString("<li class=\"active\">")
		} else {
			buffer.WriteString("<li>")
		}

		if i == 0 {
			buffer.WriteString(fmt.Sprintf("<a href=\"javascript:%s(%d, %d);\">", fn, i, size))
			buffer.WriteString("<span aria-hidden=\"true\">&laquo;</span>")
			buffer.WriteString("</a>")
			buffer.WriteString("</li>")
			continue
		}

		if i+1 == pages {
			buffer.WriteString(fmt.Sprintf("<a href=\"javascript:%s(%d, %d);\">", fn, i, size))
			buffer.WriteString("<span aria-hidden=\"true\">&raquo;</span>")
			buffer.WriteString("</a>")
			buffer.WriteString("</li>")
			continue
		}

		buffer.WriteString(fmt.Sprintf("<a href=\"javascript:%s(%d, %d);\">%d</a>", fn, i, size, i+1))
		buffer.WriteString("</li>")
	}
	buffer.WriteString("</ul>")
	buffer.WriteString("</nav>")
	return template.HTML(buffer.String())
}
