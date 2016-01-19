package controllers

import (
	models "github.com/liujianping/scaffold/demo/app/models"
	routes "github.com/liujianping/scaffold/demo/app/routes"
	"github.com/revel/revel"
)

type UserAccountController struct {
	DBController
}

func (c UserAccountController) Index() revel.Result {
	revel.TRACE.Printf("GET >> user.account.index ...")

	total, items, err := models.DefaultUserAccount.Search(db,
		models.DefaultUserAccountQuery,
		models.DefaultUserAccountSortBy,
		models.DefaultUserAccountPage)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("user.account/index.html")
	}

	pagination := Pagination("pagination", "user.account",
		total,
		models.DefaultUserAccountPage.Size,
		models.DefaultUserAccountPage.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("user.account/index.html")
}

func (c UserAccountController) Query(query models.UserAccountQuery,
	sort models.UserAccountSortBy,
	page models.UserAccountPage) revel.Result {
	revel.TRACE.Printf("POST >> user.account.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.DefaultUserAccount.Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("user.account/query.html")
	}

	pagination := Pagination("pagination", "user.account", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("user.account/query.html")
}

func (c UserAccountController) Detail(id int64) revel.Result {
	revel.TRACE.Printf("GET >> user.account.detail ... (%d)", id)

	obj, err := db.Get(models.UserAccount{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserAccountController.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("user.account/detail.html")
}

func (c UserAccountController) Create() revel.Result {
	revel.TRACE.Printf("GET >> user.account.create ...")
	return c.RenderTemplate("user.account/create.html")
}

func (c UserAccountController) CreatePost(obj models.UserAccount) revel.Result {
	revel.TRACE.Printf("POST >> user.account.create create ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(UserAccountController.Create)
	}

	c.Begin()
	if err := c.Txn.Insert(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserAccountController.Index())
	}
	c.Commit()

	c.Flash.Success("user.account (%d) create succeed.", obj.ID)
	return c.Redirect(routes.UserAccountController.Index())
}

func (c UserAccountController) Update(id int64) revel.Result {
	revel.TRACE.Printf("GET >> user.account.update ... (%d)", id)

	obj, err := db.Get(models.UserAccount{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserAccountController.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("user.account/update.html")
}

func (c UserAccountController) UpdatePost(obj models.UserAccount) revel.Result {
	revel.TRACE.Printf("POST >> user.account.update ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.UserAccountController.Update(obj.ID))
	}

	c.Begin()
	if _, err := c.Txn.Update(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserAccountController.Index())
	}
	c.Commit()

	c.Flash.Success("user.account (%d) updated success.", obj.ID)
	return c.Redirect(routes.UserAccountController.Index())
}

func (c UserAccountController) Remove(id int64) revel.Result {
	revel.TRACE.Printf("GET >> user.account.remove ... (%d)", id)

	c.Begin()
	var obj models.UserAccount
	obj.ID = id

	count, err := c.Txn.Delete(&obj)
	if err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserAccountController.Index())
	}
	c.Commit()

	if count == 0 {
		c.Flash.Error("user.account (%d) not exist.", id)
	} else {
		c.Flash.Success("user.account (%d) remove succeed.", id)
	}

	return c.Redirect(routes.UserAccountController.Index())
}

func (c UserAccountController) FinderIndex() revel.Result {
	revel.TRACE.Printf("GET >> user.account.finder.index ...")

	total, items, err := models.DefaultUserAccount.Search(db,
		models.DefaultUserAccountQuery,
		models.DefaultUserAccountSortBy,
		models.DefaultUserAccountPage)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("user.account/finder.index.html")
	}

	pagination := Pagination("finder.pagination", "user.account", total,
		models.DefaultUserAccountPage.Size,
		models.DefaultUserAccountPage.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("user.account/finder.index.html")
}

func (c UserAccountController) FinderQuery(query models.UserAccountQuery,
	sort models.UserAccountSortBy,
	page models.UserAccountPage) revel.Result {
	revel.TRACE.Printf("POST >> user.account.finder.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.DefaultUserAccount.Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("user.account/finder.query.html")
	}

	pagination := Pagination("finder.pagination", "user.account", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("user.account/finder.query.html")
}

func UserAccountCount() int64 {
	return models.DefaultUserAccount.Count(db)
}

func init() {
	revel.TemplateFuncs["UserAccountCount"] = UserAccountCount
}
