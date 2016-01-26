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

	total, items, err := models.Default[[.ClassName]].Search(db,
		models.Default[[.ClassName]]Query,
		models.Default[[.ClassName]]SortBy,
		models.Default[[.ClassName]]Page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("[[.ModuleName]]/index.html")
	}

	pagination := Pagination("pagination", "[[.ModuleName]]", 
		total,
		models.Default[[.ClassName]]Page.Size,
		models.Default[[.ClassName]]Page.No)

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

	total, items, err := models.Default[[.ClassName]].Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("[[.ModuleName]]/query.html")
	}

	pagination := Pagination("pagination", "[[.ModuleName]]", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("[[.ModuleName]]/query.html")
}

func (c [[.ClassName]]Controller) Detail(id int64) revel.Result {
	revel.TRACE.Printf("GET >> [[.ModuleName]].detail ... (%d)", id)

	obj, err := db.Get(models.[[.ClassName]]{}, id)
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

	obj, err := db.Get(models.[[.ClassName]]{}, id)
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

func (c [[.ClassName]]Controller) FinderIndex() revel.Result {
	revel.TRACE.Printf("GET >> [[.ModuleName]].finder.index ...")

	total, items, err := models.Default[[.ClassName]].Search(db,
		models.Default[[.ClassName]]Query,
		models.Default[[.ClassName]]SortBy,
		models.Default[[.ClassName]]Page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("[[.ModuleName]]/finder.index.html")
	}

	pagination := Pagination("finder.pagination", "[[.ModuleName]]", total,
		models.Default[[.ClassName]]Page.Size,
		models.Default[[.ClassName]]Page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("[[.ModuleName]]/finder.index.html")
}

func (c [[.ClassName]]Controller) FinderQuery(query models.[[.ClassName]]Query,
	sort models.[[.ClassName]]SortBy,
	page models.[[.ClassName]]Page) revel.Result {
	revel.TRACE.Printf("POST >> [[.ModuleName]].finder.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.Default[[.ClassName]].Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("[[.ModuleName]]/finder.query.html")
	}

	pagination := Pagination("finder.pagination", "[[.ModuleName]]", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("[[.ModuleName]]/finder.query.html")
}

[[if eq .ClassName "Option"]]

var selections = map[string]interface{}{}

func GetSelection(code string) []models.[[.ClassName]] {
	if items, ok := selections[code]; ok {
		return items.([]models.[[.ClassName]])
	}

	query := models.[[.ClassName]]Query{Code: code}
	sort := models.[[.ClassName]]SortBy{1}
	page := models.[[.ClassName]]Page{0, 0}
	_, items, _ := models.Default[[.ClassName]].Search(db,
		query, sort, page)

	selections[code] = items
	return items
}

func GetOption(code, option_code string) int64 {
	opts := GetSelection(code)
	for _, opt := range opts {
		if strings.ToLower(opt.OptionCode) == strings.ToLower(option_code) {
			return opt.OptionValue
		}
	}
	return 0
}

func GetOptionName(code string, option_value int64) string {
	opts := GetSelection(code)
	for _, opt := range opts {
		if opt.OptionValue == option_value {
			return opt.OptionName
		}
	}
	return "未知"
}

func init() {
	revel.TemplateFuncs["selection"] = GetSelection
	revel.TemplateFuncs["option"] = GetOption
	revel.TemplateFuncs["option_name"] = GetOptionName
}
[[end]]

func [[.ClassName]]Count() int64 {
	return models.Default[[.ClassName]].Count(db)
}

func init() {
	revel.TemplateFuncs["[[.ClassName]]Count"] = [[.ClassName]]Count
}
