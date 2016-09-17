package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
)

type LoginController struct {
    beego.Controller
}


func (c *LoginController) Get() {
    
    c.TplName = "login.html"
}


type LoginPostController struct {
    beego.Controller
}

func (c *LoginPostController) Post() {
    username := c.Input().Get("username")
    pwd := c.Input().Get("pwd")
    logs.Debug(username, pwd)
    
    if username == "golang" && pwd == "123456" {
        c.SetSession("username", username)
        c.SetSession("userid", 1)
        //c.Ctx.WriteString("登录成功")
        c.Redirect("/auth/index", 200)
    } else {
        c.Ctx.WriteString("登录失败")
    }
}

type AuthIndexController struct {
    beego.Controller
}

func (c *AuthIndexController) Get() {
    sess := c.StartSession()
    //defer sess.SessionRelease()
    
    username := sess.Get("username")
    userid := sess.Get("userid")
    
    logs.Debug(username, userid)
    
    c.Data["username"] = username
    c.Data["userid"] = userid
    
    c.TplName = "authIndex.html"
}