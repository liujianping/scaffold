package controllers
[[set . "t_class" (.table.Name | singular | camel)]]
[[set . "t_module" (.table.Name | module)]]
import (
	models "[[.project]]/app/models"
	routes "[[.project]]/app/routes"
	"github.com/revel/revel"
)

type [[.t_class]]Controller struct {
	PortalController
}

func (c [[.t_class]]Controller) Index() revel.Result {
	revel.TRACE.Printf("GET >> [[.t_module]].index ...")

	total, items, err := models.Default[[.t_class]].Search(db,
		models.Default[[.t_class]]Query,
		models.Default[[.t_class]]Sort,
		models.Default[[.t_class]]Page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("view.[[.t_module]]/index.html")
	}

	pagination := Pagination("pagination", "[[.t_module]]",
		total,
		models.Default[[.t_class]]Page.Size,
		models.Default[[.t_class]]Page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("view.[[.t_module]]/index.html")
}

func (c [[.t_class]]Controller) Query(query models.[[.t_class]]Query,
	sort models.[[.t_class]]Sort,
	page models.[[.t_class]]Page) revel.Result {
	revel.TRACE.Printf("POST >> [[.t_module]].query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	if sort.Value == 0 {
		sort.Value = models.Default[[.t_class]]Sort.Value
	}
	if page.Size == 0 {
		page.Size = models.Default[[.t_class]]Page.Size
	}

	total, items, err := models.Default[[.t_class]].Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("view.[[.t_module]]/query.html")
	}

	pagination := Pagination("pagination", "[[.t_module]]", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("view.[[.t_module]]/query.html")
}

func (c [[.t_class]]Controller) Detail(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.t_module]].detail ... (%d)", id)

	var obj models.[[.t_class]]
	db.First(&obj, id)
	if err := db.Error; err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.t_class]]Controller.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("view.[[.t_module]]/detail.html")
}

func (c [[.t_class]]Controller) Create() revel.Result {
	revel.TRACE.Printf("GET >> [[.t_module]].create ...")
	return c.RenderTemplate("view.[[.t_module]]/create.html")
}

func (c [[.t_class]]Controller) CreatePost(obj models.[[.t_class]]) revel.Result {
	revel.TRACE.Printf("POST >> [[.t_module]].create create ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect([[.t_class]]Controller.Create)
	}

	tx := db.Begin()
	tx.Create(&obj)
	if err := tx.Error; err != nil {
		tx.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.t_class]]Controller.Index())
	}
	tx.Commit()

	if err := db.Error; err != nil {
		c.Flash.Error("[[.t_module]] (%d) create failed: (%s)", obj.ID, err.Error())
	} else {
		c.Flash.Success("[[.t_module]] (%d) create succeed.", obj.ID)
	}
	return c.Redirect(routes.[[.t_class]]Controller.Index())
}

func (c [[.t_class]]Controller) Update(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.t_module]].update ... (%d)", id)

	var obj models.[[.t_class]]
	db.First(&obj, id)
	if err := db.Error; err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.t_class]]Controller.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("view.[[.t_module]]/update.html")
}

func (c [[.t_class]]Controller) UpdatePost(obj models.[[.t_class]]) revel.Result {
	revel.TRACE.Printf("POST >> [[.t_module]].update ... (%v)", obj)

	tx := db.Begin()
	tx.Model(&obj).Updates(obj)
	tx.Commit()

	if err := db.Error; err != nil {
		c.Flash.Error("[[.t_module]] (%d) update failed: (%s)", obj.ID, err.Error())
	} else {
		c.Flash.Success("[[.t_module]] (%d) update succeed.", obj.ID)
	}
	return c.Redirect(routes.[[.t_class]]Controller.Index())
}

func (c [[.t_class]]Controller) Remove(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.t_module]].remove ... (%d)", id)

	var obj models.[[.t_class]]
	obj.ID = id

	tx := db.Begin()
	tx.Delete(&obj)
	tx.Commit()

	if err := db.Error; err != nil {
		c.Flash.Error("[[.t_module]] (%d) remove failed: (%s)", id, err.Error())
	} else {
		c.Flash.Success("[[.t_module]] (%d) remove succeed.", id)
	}
	return c.Redirect(routes.[[.t_class]]Controller.Index())
}

func (c [[.t_class]]Controller) RemovePost(query models.[[.t_class]]Query,
	sort models.[[.t_class]]Sort,
	page models.[[.t_class]]Page,
	id []int64) revel.Result {
	revel.TRACE.Printf("POST >> [[.t_module]].remove ...(query: %v) (sort: %v) (page: %v) (id: %v)",
		query, sort, page, id)

	tx := db.Begin()
	tx.Where(id).Delete(models.Default[[.t_class]])
	tx.Commit()

	if err := db.Error; err != nil {
		c.Flash.Error("[[.t_module]] (%v) remove failed: (%s)", id, err.Error())
	} else {
		c.Flash.Success("[[.t_module]] (%v) remove succeed.", id)
	}
	return c.Query(query, sort, page)
}

func (c [[.t_class]]Controller) FinderIndex() revel.Result {
	revel.TRACE.Printf("GET >> [[.t_module]].finder.index ...")

	total, items, err := models.Default[[.t_class]].Search(db,
		models.Default[[.t_class]]Query,
		models.Default[[.t_class]]Sort,
		models.Default[[.t_class]]Page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("view.[[.t_module]]/finder.index.html")
	}

	pagination := Pagination("finder.pagination", "[[.t_module]]", total,
		models.Default[[.t_class]]Page.Size,
		models.Default[[.t_class]]Page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("view.[[.t_module]]/finder.index.html")
}

func (c [[.t_class]]Controller) FinderQuery(query models.[[.t_class]]Query,
	sort models.[[.t_class]]Sort,
	page models.[[.t_class]]Page) revel.Result {
	revel.TRACE.Printf("POST >> [[.t_module]].finder.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.Default[[.t_class]].Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("view.[[.t_module]]/finder.query.html")
	}

	pagination := Pagination("finder.pagination", "[[.t_module]]", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("view.[[.t_module]]/finder.query.html")
}

[[if eq (.table.Tag "import") "y"]]
func (c [[.t_class]]Controller) Import() revel.Result {
	revel.INFO.Printf("POST >> [[.t_module]].import ...")
	return c.Redirect(routes.[[.t_class]]Controller.Index())
}
[[end]]

[[if eq (.table.Tag "export") "y"]]
func (c [[.t_class]]Controller) Export() revel.Result {
	revel.INFO.Printf("POST >> [[.t_module]].export ...")
	return c.Redirect(routes.[[.t_class]]Controller.Index())
}
[[end]]

func [[.t_class]]Count() int64 {
	return models.Default[[.t_class]].Count(db)
}

func init() {
	revel.TemplateFuncs["[[.t_class]]Count"] = [[.t_class]]Count
}
