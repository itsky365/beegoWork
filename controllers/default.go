package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sess := c.StartSession()

	username := sess.Get("username")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["username"] = username

	c.TplName = "index.html"
	//c.Redirect("/login", 302)
	//c.Ctx.Redirect(302, "/login")
}
