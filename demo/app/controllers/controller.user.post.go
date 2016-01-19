package controllers

import (
	models "github.com/liujianping/scaffold/demo/app/models"
	routes "github.com/liujianping/scaffold/demo/app/routes"
	"github.com/revel/revel"
)

type UserPostController struct {
	DBController
}

func (c UserPostController) Index() revel.Result {
	revel.TRACE.Printf("GET >> user.post.index ...")

	total, items, err := models.DefaultUserPost.Search(db,
		models.DefaultUserPostQuery,
		models.DefaultUserPostSortBy,
		models.DefaultUserPostPage)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("user.post/index.html")
	}

	pagination := Pagination("pagination", "user.post",
		total,
		models.DefaultUserPostPage.Size,
		models.DefaultUserPostPage.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("user.post/index.html")
}

func (c UserPostController) Query(query models.UserPostQuery,
	sort models.UserPostSortBy,
	page models.UserPostPage) revel.Result {
	revel.TRACE.Printf("POST >> user.post.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.DefaultUserPost.Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("user.post/query.html")
	}

	pagination := Pagination("pagination", "user.post", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("user.post/query.html")
}

func (c UserPostController) Detail(id int64) revel.Result {
	revel.TRACE.Printf("GET >> user.post.detail ... (%d)", id)

	obj, err := db.Get(models.UserPost{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserPostController.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("user.post/detail.html")
}

func (c UserPostController) Create() revel.Result {
	revel.TRACE.Printf("GET >> user.post.create ...")
	return c.RenderTemplate("user.post/create.html")
}

func (c UserPostController) CreatePost(obj models.UserPost) revel.Result {
	revel.TRACE.Printf("POST >> user.post.create create ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(UserPostController.Create)
	}

	c.Begin()
	if err := c.Txn.Insert(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserPostController.Index())
	}
	c.Commit()

	c.Flash.Success("user.post (%d) create succeed.", obj.ID)
	return c.Redirect(routes.UserPostController.Index())
}

func (c UserPostController) Update(id int64) revel.Result {
	revel.TRACE.Printf("GET >> user.post.update ... (%d)", id)

	obj, err := db.Get(models.UserPost{}, id)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserPostController.Index())
	}

	c.RenderArgs["obj"] = obj
	return c.RenderTemplate("user.post/update.html")
}

func (c UserPostController) UpdatePost(obj models.UserPost) revel.Result {
	revel.TRACE.Printf("POST >> user.post.update ... (%v)", obj)

	obj.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.UserPostController.Update(obj.ID))
	}

	c.Begin()
	if _, err := c.Txn.Update(&obj); err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserPostController.Index())
	}
	c.Commit()

	c.Flash.Success("user.post (%d) updated success.", obj.ID)
	return c.Redirect(routes.UserPostController.Index())
}

func (c UserPostController) Remove(id int64) revel.Result {
	revel.TRACE.Printf("GET >> user.post.remove ... (%d)", id)

	c.Begin()
	var obj models.UserPost
	obj.ID = id

	count, err := c.Txn.Delete(&obj)
	if err != nil {
		c.Rollback()
		c.Flash.Error(err.Error())
		return c.Redirect(routes.UserPostController.Index())
	}
	c.Commit()

	if count == 0 {
		c.Flash.Error("user.post (%d) not exist.", id)
	} else {
		c.Flash.Success("user.post (%d) remove succeed.", id)
	}

	return c.Redirect(routes.UserPostController.Index())
}

func (c UserPostController) FinderIndex() revel.Result {
	revel.TRACE.Printf("GET >> user.post.finder.index ...")

	total, items, err := models.DefaultUserPost.Search(db,
		models.DefaultUserPostQuery,
		models.DefaultUserPostSortBy,
		models.DefaultUserPostPage)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("user.post/finder.index.html")
	}

	pagination := Pagination("finder.pagination", "user.post", total,
		models.DefaultUserPostPage.Size,
		models.DefaultUserPostPage.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("user.post/finder.index.html")
}

func (c UserPostController) FinderQuery(query models.UserPostQuery,
	sort models.UserPostSortBy,
	page models.UserPostPage) revel.Result {
	revel.TRACE.Printf("POST >> user.post.finder.query ...(query: %v) (sort: %v) (page: %v)",
		query, sort, page)

	total, items, err := models.DefaultUserPost.Search(db,
		query, sort, page)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.RenderTemplate("user.post/finder.query.html")
	}

	pagination := Pagination("finder.pagination", "user.post", total, page.Size, page.No)

	c.RenderArgs["total"] = total
	c.RenderArgs["items"] = items
	c.RenderArgs["pagination"] = pagination
	return c.RenderTemplate("user.post/finder.query.html")
}

func UserPostCount() int64 {
	return models.DefaultUserPost.Count(db)
}

func init() {
	revel.TemplateFuncs["UserPostCount"] = UserPostCount
}
