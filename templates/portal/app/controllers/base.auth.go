package controllers

import (
	"[[.project]]/app/routes"
	"github.com/revel/revel"
)

type AuthController struct {
	*revel.Controller
}

func (c AuthController) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}

	return c.Redirect(routes.AuthController.Login())
}

func (c AuthController) Password() revel.Result {
	// if account := c.connected(); account == nil {
	// 	c.Flash.Error("please login first.")
	// 	return c.Redirect(routes.AuthController.Login())
	// }

	return c.RenderTemplate("auth/password.html")
}

func (c AuthController) PasswordPost() revel.Result {
	revel.INFO.Printf("POST > /auth.password ...")
	var old_password, new_password, new_password2 string

	c.Params.Bind(&old_password, "old_password")
	c.Params.Bind(&new_password, "new_password")
	c.Params.Bind(&new_password2, "new_password2")

	if new_password != new_password2 {
		c.Flash.Error("新密码不一致。")
		return c.Redirect(routes.AuthController.Password())
	}

	//! params validation check
	c.Validation.Required(old_password).Message("原密码不能为空")
	c.Validation.Required(new_password).Message("新密码不能为空")

	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.AuthController.Password())
	}

	// err := business.PortalAccountPassword(db, account, old_password, new_password)
	// if err != nil {
	// 	c.Flash.Error("reset password:" + err.Error())
	// 	return c.Redirect(routes.AuthController.Login())
	// }

	c.Flash.Success("auth.password updated success.")
	return c.Redirect(routes.PortalController.Index())
}

func (c AuthController) Login() revel.Result {
	return c.RenderTemplate("auth/login.html")
}

func (c AuthController) LoginPost() revel.Result {
	revel.INFO.Printf("POST > /login ...")
	var name, password string
	var remember bool

	c.Params.Bind(&name, "name")
	c.Params.Bind(&password, "password")
	c.Params.Bind(&remember, "remember")

	revel.INFO.Println("name =>", name)
	revel.INFO.Println("password =>", password)
	//! params validation check
	c.Validation.Required(name).Message("账号邮箱不能为空")
	c.Validation.Email(name).Message("请输入邮箱地址")
	c.Validation.Required(password).Message("账号密码不能为空")
	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.AuthController.Login())
	}

	// account, err := business.AuthorizePortalAccount(db, name, password)
	// if err != nil {
	// 	c.Flash.Error("authorized failed:" + err.Error())
	// 	return c.Redirect(routes.AuthController.Login())
	// }

	// if account.Status != models.GeneralStatusNormal {
	// 	c.Flash.Error("authorization forbidden")
	// 	return c.Redirect(routes.AuthController.Login())
	// }

	// expires := revel.Config.StringDefault("session.expires", "10m")
	// d, err := time.ParseDuration(expires)
	// if err != nil {
	// 	c.Flash.Error("time parsed failed:" + err.Error())
	// 	return c.Redirect(routes.AuthController.Login())
	// }

	// portal := models.AuthCategoryPortal

	// remote, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	// forward := c.Request.Header.Get("X-Forwarded-For")
	// if forward != "" {
	// 	remote = forward
	// }

	// token, err := business.AuthTokenGen(db, account.ID, portal, account.Role, remote, 0, d)
	// if err != nil {
	// 	c.Flash.Error("token generate failed:" + err.Error())
	// 	return c.Redirect(routes.AuthController.Login())
	// }

	c.Session["secret"] = "authorized"
	if remember {
		c.Session.SetDefaultExpiration()
	} else {
		c.Session.SetNoExpiration()
	}

	return c.Redirect(routes.PortalController.Index())
}
