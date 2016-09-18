package routers

import (
    "beegoWork/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/blog/:id:int", &controllers.BlogIndexController{})
    beego.Router("/blog/json", &controllers.BlogJsonController{})
    beego.Router("/blog/add", &controllers.BlogAddController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/logout", &controllers.LogoutController{})
    beego.Router("/auth/index", &controllers.AuthIndexController{})
}
