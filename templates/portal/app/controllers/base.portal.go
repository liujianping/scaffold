package controllers

import (
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
		return secret, true
		// remote, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		// forward := c.Request.Header.Get("X-Forwarded-For")
		// if forward != "" {
		// 	remote = forward
		// }
		// token, err := business.AuthTokenHit(db, secret, remote)
		// if err != nil {
		// 	revel.TRACE.Printf("token (%s) hit failed:(%s)", secret, err.Error())
		// 	return nil
		// }

		// if token == nil {
		// 	revel.TRACE.Printf("token (%s) hit invalid.", secret)
		// 	return nil
		// }

		// obj, err := db.Get(models.PortalAccount{}, token.AuthID)
		// if err != nil {
		// 	return nil
		// }

		// c.RenderArgs["account"] = obj.(*models.PortalAccount)
		// return obj.(*models.PortalAccount)
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
