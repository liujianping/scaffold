package controllers

import (
	"net"

	"[[.project]]/app/models"
	"[[.project]]/app/routes"
	"github.com/revel/revel"
)

type PortalController struct {
	*revel.Controller
}

func (c PortalController) Index() revel.Result {
	return c.RenderTemplate("index.html")
}

func (c PortalController) authorized() (interface{}, bool) {
	if secret, ok := c.Session["secret"]; ok {
		remote, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		forward := c.Request.Header.Get("X-Forwarded-For")
		if forward != "" {
			remote = forward
		}

		token, err := HitSystemToken(secret, remote)
		if err != nil {
			revel.WARN.Printf("authorized failed: %s", err.Error())
			return nil, false
		}

		var account models.SystemAccount
		account.ID = token.SystemAccountID
		if err := db.First(&account).Error; err != nil {
			return nil, false
		}
		c.RenderArgs["account"] = &account
		return &token, true
	}

	return nil, false
}

func (c PortalController) Authorized() revel.Result {
	if _, ok := c.authorized(); !ok {
		return c.Redirect(routes.AuthController.Login())
	}
	return nil
}

func init() {
	revel.InterceptMethod(PortalController.Authorized, revel.BEFORE)
}
