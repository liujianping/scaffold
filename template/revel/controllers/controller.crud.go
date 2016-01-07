package controllers
[[set . "ClassName" (.table.Name | singular | camel)]]
[[set . "ModuleName" (.table.Name | module)]]
import (
	"github.com/revel/revel"
	models "[[.project]]/app/models"
	routes "[[.project]]/app/routes"
)

type [[.ClassName]]Controller struct {
	DBController
}

func (c [[.ClassName]]Controller) Index() revel.Result {
	revel.TRACE.Printf("GET >> [[.ModuleName]].index ...")

	total, items, err := models.Default[[.ClassName]].Execute(model,
		models.Default[[.ClassName]]Query,
		models.Default[[.ClassName]]SortBy,
		models.Default[[.ClassName]]Page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderError(err)
	}

	pagination := c.Pagination("search",
		int(total),
		models.Default[[.ClassName]]Page.No,
		models.PageSize(models.Default[[.ClassName]]Page.Size))

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("[[.ModuleName]]/index.html")
}

func (c [[.ClassName]]Controller) Query(query models.[[.ClassName]]Query,
	sort models.[[.ClassName]]SortBy,
	page models.[[.ClassName]]Page) revel.Result {
	revel.TRACE.Printf("POST >> [[.ModuleName]].query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.Default[[.ClassName]].Execute(model,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("[[.ModuleName]]/query.html")
	}

	pagination := c.Pagination("search",
		int(total),
		page.No,
		models.PageSize(page.Size))

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("[[.ModuleName]]/query.html")
}

func (c [[.ClassName]]Controller) Detail(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.ModuleName]].detail ... (%d)", id)

	obj, err := model.Get(models.[[.ClassName]]{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.ClassName]]Controller.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("[[.ModuleName]]/detail.html")
}

func (c [[.ClassName]]Controller) Create() revel.Result {
	revel.TRACE.Printf("GET >> [[.ModuleName]].create ...")
	return c.RenderTemplate("[[.ModuleName]]/create.html")
}

func (c [[.ClassName]]Controller) CreatePost(obj models.[[.ClassName]]) revel.Result {
	revel.TRACE.Printf("POST >> [[.ModuleName]].create create ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect([[.ClassName]]Controller.Create)
	}

	c.Begin()
	if err := c.Txn.Insert(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.ClassName]]Controller.Index())
	}
	c.Commit()

	c.Flash.Success("[[.ModuleName]] (%d) create succeed.", obj.ID)
	return c.Redirect(routes.[[.ClassName]]Controller.Index())
}

func (c [[.ClassName]]Controller) Update(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.ModuleName]].update ... (%d)", id)

	obj, err := model.Get(models.[[.ClassName]]{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.ClassName]]Controller.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("[[.ModuleName]]/update.html")
}

func (c [[.ClassName]]Controller) UpdatePost(obj models.[[.ClassName]]) revel.Result {
	revel.TRACE.Printf("POST >> [[.ModuleName]].update ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.[[.ClassName]]Controller.Update(obj.ID))
	}

	c.Begin()
	if _, err := c.Txn.Update(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.ClassName]]Controller.Index())
	}
	c.Commit()

	c.Flash.Success("[[.ModuleName]] (%d) updated success.", obj.ID)
	return c.Redirect(routes.[[.ClassName]]Controller.Index())
}

func (c [[.ClassName]]Controller) Remove(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.ModuleName]].remove ... (%d)", id)

	c.Begin()
	var obj models.[[.ClassName]]
	obj.ID = id

	count, err := c.Txn.Delete(&obj)
	if err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.ClassName]]Controller.Index())
	}
	c.Commit()

	if count == 0 {
		c.Flash.Error("[[.ModuleName]] (%d) not exist.", id)
	} else {
		c.Flash.Success("[[.ModuleName]] (%d) remove succeed.", id)
	}

	return c.Redirect(routes.[[.ClassName]]Controller.Index())
}
