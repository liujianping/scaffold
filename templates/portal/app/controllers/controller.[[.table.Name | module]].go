package controllers
[[set . "t_class" (.table.Name | singular | camel)]]
[[set . "t_module" (.table.Name | module)]]
import (
	"github.com/gocarina/gocsv"
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
	if err := db.First(&obj, id).Error; err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.[[.t_class]]Controller.Index())
	}

	[[range .tables]]
	[[if eq (.Tag "belong") $.table.Name]]    
	[[if ne (.Tag "many") ""]]
	db.Model(&obj).Association("[[.Tag "many" | camel | lint]]").Find(&obj.[[.Tag "many" | camel | lint]])
	[[end]][[end]][[end]]

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
	if err := tx.Create(models.[[.t_class]]Cipher(&obj)).Error; err != nil {
		tx.Rollback()
		c.Flash.Error("[[.t_module]] (%d) create failed: (%s)", obj.ID, err.Error())
		return c.Redirect(routes.[[.t_class]]Controller.Index())
	}
	tx.Commit()

	c.Flash.Success("[[.t_module]] (%d) create succeed.", obj.ID)
	return c.Redirect(routes.[[.t_class]]Controller.Index())
}

func (c [[.t_class]]Controller) Update(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.t_module]].update ... (%d)", id)

	var obj models.[[.t_class]]
	if err := db.First(&obj, id).Error; err != nil {
		c.Flash.Error("[[.t_module]] (%d) update failed: (%s)", id, err.Error())
		return c.Redirect(routes.[[.t_class]]Controller.Index())
	}

	[[range .tables]]
	[[if eq (.Tag "belong") $.table.Name]]    
	[[if ne (.Tag "many") ""]]
	db.Model(&obj).Association("[[.Tag "many" | camel | lint]]").Find(&obj.[[.Tag "many" | camel | lint]])
	[[end]][[end]][[end]]

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("view.[[.t_module]]/update.html")
}

func (c [[.t_class]]Controller) UpdatePost(obj models.[[.t_class]]) revel.Result {
	revel.TRACE.Printf("POST >> [[.t_module]].update ... (%v)", obj)

	tx := db.Begin()
	if err := tx.Model(&obj).Updates(models.[[.t_class]]Cipher(&obj)).Error; err != nil {
		tx.Rollback()
		c.Flash.Error("[[.t_module]] (%d) update failed: (%s)", obj.ID, err.Error())
		return c.Redirect(routes.[[.t_class]]Controller.Index())
	}
	tx.Commit()

	c.Flash.Success("[[.t_module]] (%d) update succeed.", obj.ID)
	return c.Redirect(routes.[[.t_class]]Controller.Index())
}

func (c [[.t_class]]Controller) Remove(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.t_module]].remove ... (%d)", id)

	var obj models.[[.t_class]]
	obj.ID = id

	tx := db.Begin()
	if err := tx.Delete(&obj).Error; err != nil {
		tx.Rollback()
		c.Flash.Error("[[.t_module]] (%d) remove failed: (%s)", id, err.Error())
		return c.Redirect(routes.[[.t_class]]Controller.Index())
	}
	tx.Commit()

	c.Flash.Success("[[.t_module]] (%d) remove succeed.", id)
	return c.Redirect(routes.[[.t_class]]Controller.Index())
}

func (c [[.t_class]]Controller) RemovePost(query models.[[.t_class]]Query,
	sort models.[[.t_class]]Sort,
	page models.[[.t_class]]Page,
	id []int64) revel.Result {
	revel.TRACE.Printf("POST >> [[.t_module]].remove ...(query: %v) (sort: %v) (page: %v) (id: %v)",
		query, sort, page, id)

	tx := db.Begin()
	if err := tx.Where(id).Delete(models.Default[[.t_class]]).Error; err != nil {
		tx.Rollback()
		c.Flash.Error("[[.t_module]] (%v) remove failed: (%s)", id, err.Error())
		return c.Query(query, sort, page)
	}
	tx.Commit()

	c.Flash.Success("[[.t_module]] (%v) remove succeed.", id)
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
func (c [[.t_class]]Controller) ImportPost() revel.Result {
	revel.INFO.Printf("POST >> [[.t_module]].import ...")
	
	c.Validation.Required(c.Params.Files["upload"])
	if c.Validation.HasErrors() {
		return c.RenderJson(WidgetResponse{Code: 401, Message: "files absent"})
	}

	upload_file := c.Params.Files["upload"][0]
	fd, err := upload_file.Open()
	if err != nil {
		return c.RenderJson(WidgetResponse{Code: 402, Message: err.Error()})
	}
	defer fd.Close()

	items := []*models.[[.t_class]]{}
	if err := gocsv.Unmarshal(fd, &items); err != nil {
		return c.RenderJson(WidgetResponse{Code: 403, Message: err.Error()})
	}

	tx := db.Begin()
	for _, item := range items {
		if err := tx.Create(item).Error; err != nil {
			tx.Rollback()
			return c.RenderJson(WidgetResponse{Code: 405, Message: err.Error()})
		}
	}
	tx.Commit()

	return c.RenderJson(WidgetResponse{Code: 0, Message: fmt.Sprintf("[[.t_module]] %d imported", len(items))})
}
[[end]]

[[if eq (.table.Tag "export") "y"]]
func (c [[.t_class]]Controller) ExportPost(query models.[[.t_class]]Query) revel.Result {
	revel.INFO.Printf("POST >> [[.t_module]].export ...")

	items, err := models.Default[[.t_class]].Export(db, query)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("view.[[.t_module]]/query.html")
	}

	bio := bytes.NewBuffer([]byte{})
	if err := gocsv.Marshal(items, bio); err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("view.[[.t_module]]/query.html")
	}

	url, err := Upload(fmt.Sprintf("[[.t_module]].%s.csv", time.Now().Format("20060102150405")), bio)
	if err != nil {
		return c.RenderJson(map[string]interface{}{
			"code":  405,
			"error": err.Error(),
		})
	}

	return c.RenderJson(map[string]interface{}{
		"code": 0,
		"url":  url,
	})
}
[[end]]

func [[.t_class]]Count() int64 {
	return models.Default[[.t_class]].Count(db)
}

func [[.t_class]]Field(field string, id int64) interface{} {
	return models.[[.t_class]]Field(db, field, id)
}

func init() {
	revel.TemplateFuncs["[[.t_class]]Count"] = [[.t_class]]Count
	revel.TemplateFuncs["[[.t_class]]Field"] = [[.t_class]]Field
}
