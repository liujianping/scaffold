package controllers

import (
	"strings"

	models "github.com/liujianping/scaffold/demo/app/models"
	routes "github.com/liujianping/scaffold/demo/app/routes"
	"github.com/revel/revel"
)

type OptionController struct {
	DBController
}

func (c OptionController) Index() revel.Result {
	revel.TRACE.Printf("GET >> option.index ...")

	total, items, err := models.DefaultOption.Search(db,
		models.DefaultOptionQuery,
		models.DefaultOptionSortBy,
		models.DefaultOptionPage)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("option/index.html")
	}

	pagination := Pagination("pagination", "option",
		total,
		models.DefaultOptionPage.Size,
		models.DefaultOptionPage.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("option/index.html")
}

func (c OptionController) Query(query models.OptionQuery,
	sort models.OptionSortBy,
	page models.OptionPage) revel.Result {
	revel.TRACE.Printf("POST >> option.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.DefaultOption.Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("option/query.html")
	}

	pagination := Pagination("pagination", "option", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("option/query.html")
}

func (c OptionController) Detail(id int64) revel.Result {
	revel.TRACE.Printf("GET >> option.detail ... (%d)", id)

	obj, err := db.Get(models.Option{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.OptionController.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("option/detail.html")
}

func (c OptionController) Create() revel.Result {
	revel.TRACE.Printf("GET >> option.create ...")
	return c.RenderTemplate("option/create.html")
}

func (c OptionController) CreatePost(obj models.Option) revel.Result {
	revel.TRACE.Printf("POST >> option.create create ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(OptionController.Create)
	}

	c.Begin()
	if err := c.Txn.Insert(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.OptionController.Index())
	}
	c.Commit()

	c.Flash.Success("option (%d) create succeed.", obj.ID)
	return c.Redirect(routes.OptionController.Index())
}

func (c OptionController) Update(id int64) revel.Result {
	revel.TRACE.Printf("GET >> option.update ... (%d)", id)

	obj, err := db.Get(models.Option{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.OptionController.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("option/update.html")
}

func (c OptionController) UpdatePost(obj models.Option) revel.Result {
	revel.TRACE.Printf("POST >> option.update ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.OptionController.Update(obj.ID))
	}

	c.Begin()
	if _, err := c.Txn.Update(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.OptionController.Index())
	}
	c.Commit()

	c.Flash.Success("option (%d) updated success.", obj.ID)
	return c.Redirect(routes.OptionController.Index())
}

func (c OptionController) Remove(id int64) revel.Result {
	revel.TRACE.Printf("GET >> option.remove ... (%d)", id)

	c.Begin()
	var obj models.Option
	obj.ID = id

	count, err := c.Txn.Delete(&obj)
	if err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.OptionController.Index())
	}
	c.Commit()

	if count == 0 {
		c.Flash.Error("option (%d) not exist.", id)
	} else {
		c.Flash.Success("option (%d) remove succeed.", id)
	}

	return c.Redirect(routes.OptionController.Index())
}

func (c OptionController) FinderIndex() revel.Result {
	revel.TRACE.Printf("GET >> option.finder.index ...")

	total, items, err := models.DefaultOption.Search(db,
		models.DefaultOptionQuery,
		models.DefaultOptionSortBy,
		models.DefaultOptionPage)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("option/finder.index.html")
	}

	pagination := Pagination("finder.pagination", "option", total,
		models.DefaultOptionPage.Size,
		models.DefaultOptionPage.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("option/finder.index.html")
}

func (c OptionController) FinderQuery(query models.OptionQuery,
	sort models.OptionSortBy,
	page models.OptionPage) revel.Result {
	revel.TRACE.Printf("POST >> option.finder.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.DefaultOption.Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("option/finder.query.html")
	}

	pagination := Pagination("finder.pagination", "option", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("option/finder.query.html")
}

var selections = map[string]interface{}{}

func GetSelection(code string) []models.Option {
	if items, ok := selections[code]; ok {
		return items.([]models.Option)
	}

	query := models.OptionQuery{Code: code}
	sort := models.OptionSortBy{1}
	page := models.OptionPage{0, 0}
	_, items, _ := models.DefaultOption.Search(db,
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

func OptionCount() int64 {
	return models.DefaultOption.Count(db)
}

func init() {
	revel.TemplateFuncs["OptionCount"] = OptionCount
}
