package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
    c.TplName = "index.html"
    
	sess := c.StartSession()

	username := sess.Get("username")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["username"] = username

	
	//c.Redirect("/login", 302)
	//c.Ctx.Redirect(302, "/login")
	
	// sets this.Data["paginator"] with the current offset (from the url query param)
	postsPerPage := 20
    pagination.SetPaginator(c.Ctx, postsPerPage, 105)
    //c.Data["paginator"] = paginator
	
	// fetch the next 20 posts
	//c.Data["posts"] = ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
}
