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
    beego.Router("/user/add", &controllers.UserAddController{})
    beego.Router("/user/edit/:id:int", &controllers.UserEditController{})
    beego.Router("/myorm", &controllers.MyOrmController{})
    beego.Router("/myupload/", &controllers.MyuploadController{})
    beego.Router("/myupload/post", &controllers.MyuploadPostController{})
}
