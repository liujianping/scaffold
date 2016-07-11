package controllers

import (
	"[[.project]]/app/models"
	"[[.project]]/app/routes"
	"github.com/revel/revel"
	"net"
	"time"
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
	if _, ok := c.Session["account"]; !ok {
		c.Flash.Error("please login first.")
		return c.Redirect(routes.AuthController.Login())
	}

	return c.RenderTemplate("auth/password.html")
}

func (c AuthController) PasswordPost() revel.Result {
	revel.INFO.Printf("POST > /auth.password ...")
	accountID, ok := c.Session["account"]
	if !ok {
		c.Flash.Error("please login first.")
		return c.Redirect(routes.AuthController.Login())
	}


	var account models.SystemAccount
	account.ID = models.DecodeID(accountID)
	if err := db.First(&account).Error; err != nil {
		c.Flash.Error("please login first: fake session.")
		return c.Redirect(routes.AuthController.Login())
	}

	var old_password, new_password, new_password2 string
	c.Params.Bind(&old_password, "old_password")
	c.Params.Bind(&new_password, "new_password")
	c.Params.Bind(&new_password2, "new_password2")

	if new_password != new_password2 {
		c.Flash.Error("new password not equal")
		return c.Redirect(routes.AuthController.Password())
	}

	//! params validation check
	c.Validation.Required(old_password)
	c.Validation.Required(new_password)

	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.AuthController.Password())
	}


	tx := db.Begin()
	account.Password = new_password
	if err := tx.Model(&account).Update("password", models.SystemAccountCipher(&account).Password).Error; err != nil {
		tx.Rollback()
		c.Flash.Error("reset password:" + err.Error())
		return c.Redirect(routes.AuthController.Login())
	}
	tx.Commit()

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

	//! params validation check
	c.Validation.Required(name)
	c.Validation.Email(name)
	c.Validation.Required(password)
	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.AuthController.Login())
	}

	var account models.SystemAccount
	if err := db.Where("mailbox = ?", name).First(&account).Error; err != nil {
		c.Flash.Error("authorized failed:" + err.Error())
		return c.Redirect(routes.AuthController.Login())
	}

	obj := models.SystemAccountCipher(&models.SystemAccount{Password: password})
	if obj.Password != account.Password {
		c.Flash.Error("authorized failed: password wrong")
		return c.Redirect(routes.AuthController.Login())
	}

	if account.Status != GetOptionValue("system_account_status", "normal") {
		c.Flash.Error("authorized failed: forbidden")
		return c.Redirect(routes.AuthController.Login())
	}

	expires := revel.Config.StringDefault("session.expires", "10m")
	duration, err := time.ParseDuration(expires)
	if err != nil {
		c.Flash.Error("time parsed failed:" + err.Error())
		return c.Redirect(routes.AuthController.Login())
	}

	remote, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	forward := c.Request.Header.Get("X-Forwarded-For")
	if forward != "" {
		remote = forward
	}

	token, err := GenerateSystemToken(account.ID, remote, 0, duration)
	if err != nil {
		c.Flash.Error("token generate failed:" + err.Error())
		return c.Redirect(routes.AuthController.Login())
	}
	c.Session["account"] = models.EncodeID(account.ID)
	c.Session["secret"] = token.Secret
	if remember {
		c.Session.SetDefaultExpiration()
	} else {
		c.Session.SetNoExpiration()
	}

	return c.Redirect(routes.PortalController.Index())
}
