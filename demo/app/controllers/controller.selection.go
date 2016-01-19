package controllers

import (
	models "github.com/liujianping/scaffold/demo/app/models"
	routes "github.com/liujianping/scaffold/demo/app/routes"
	"github.com/revel/revel"
)

type SelectionController struct {
	DBController
}

func (c SelectionController) Index() revel.Result {
	revel.TRACE.Printf("GET >> selection.index ...")

	total, items, err := models.DefaultSelection.Search(db,
		models.DefaultSelectionQuery,
		models.DefaultSelectionSortBy,
		models.DefaultSelectionPage)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("selection/index.html")
	}

	pagination := Pagination("pagination", "selection",
		total,
		models.DefaultSelectionPage.Size,
		models.DefaultSelectionPage.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("selection/index.html")
}

func (c SelectionController) Query(query models.SelectionQuery,
	sort models.SelectionSortBy,
	page models.SelectionPage) revel.Result {
	revel.TRACE.Printf("POST >> selection.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.DefaultSelection.Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("selection/query.html")
	}

	pagination := Pagination("pagination", "selection", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("selection/query.html")
}

func (c SelectionController) Detail(id int64) revel.Result {
	revel.TRACE.Printf("GET >> selection.detail ... (%d)", id)

	obj, err := db.Get(models.Selection{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.SelectionController.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("selection/detail.html")
}

func (c SelectionController) Create() revel.Result {
	revel.TRACE.Printf("GET >> selection.create ...")
	return c.RenderTemplate("selection/create.html")
}

func (c SelectionController) CreatePost(obj models.Selection) revel.Result {
	revel.TRACE.Printf("POST >> selection.create create ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(SelectionController.Create)
	}

	c.Begin()
	if err := c.Txn.Insert(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.SelectionController.Index())
	}
	c.Commit()

	c.Flash.Success("selection (%d) create succeed.", obj.ID)
	return c.Redirect(routes.SelectionController.Index())
}

func (c SelectionController) Update(id int64) revel.Result {
	revel.TRACE.Printf("GET >> selection.update ... (%d)", id)

	obj, err := db.Get(models.Selection{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.SelectionController.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("selection/update.html")
}

func (c SelectionController) UpdatePost(obj models.Selection) revel.Result {
	revel.TRACE.Printf("POST >> selection.update ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.SelectionController.Update(obj.ID))
	}

	c.Begin()
	if _, err := c.Txn.Update(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.SelectionController.Index())
	}
	c.Commit()

	c.Flash.Success("selection (%d) updated success.", obj.ID)
	return c.Redirect(routes.SelectionController.Index())
}

func (c SelectionController) Remove(id int64) revel.Result {
	revel.TRACE.Printf("GET >> selection.remove ... (%d)", id)

	c.Begin()
	var obj models.Selection
	obj.ID = id

	count, err := c.Txn.Delete(&obj)
	if err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.SelectionController.Index())
	}
	c.Commit()

	if count == 0 {
		c.Flash.Error("selection (%d) not exist.", id)
	} else {
		c.Flash.Success("selection (%d) remove succeed.", id)
	}

	return c.Redirect(routes.SelectionController.Index())
}

func (c SelectionController) FinderIndex() revel.Result {
	revel.TRACE.Printf("GET >> selection.finder.index ...")

	total, items, err := models.DefaultSelection.Search(db,
		models.DefaultSelectionQuery,
		models.DefaultSelectionSortBy,
		models.DefaultSelectionPage)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("selection/finder.index.html")
	}

	pagination := Pagination("finder.pagination", "selection", total,
		models.DefaultSelectionPage.Size,
		models.DefaultSelectionPage.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("selection/finder.index.html")
}

func (c SelectionController) FinderQuery(query models.SelectionQuery,
	sort models.SelectionSortBy,
	page models.SelectionPage) revel.Result {
	revel.TRACE.Printf("POST >> selection.finder.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.DefaultSelection.Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("selection/finder.query.html")
	}

	pagination := Pagination("finder.pagination", "selection", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("selection/finder.query.html")
}

func SelectionCount() int64 {
	return models.DefaultSelection.Count(db)
}

func init() {
	revel.TemplateFuncs["SelectionCount"] = SelectionCount
}
