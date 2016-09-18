package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego/orm"
    "beegoWork/models"
    "fmt"
)

// 用户登录
type LoginController struct {
    beego.Controller
}

func (c *LoginController) Get() {
    c.TplName = "login.html"
}

func (c *LoginController) Post() {
    //logs.Debug("username=>", c.GetString("username"))
    username := c.GetString("username")
    pwd := c.GetString("pwd")
    logs.Debug(username, pwd)
    
    if username == "golang" && pwd == "123456" {
        c.SetSession("username", username)
        c.SetSession("userid", 1)
        //c.Ctx.WriteString("登录成功")
        c.Redirect("/auth/index", 302)
    } else {
        c.Ctx.WriteString("登录失败")
    }
}


// 用户中心首页
type AuthIndexController struct {
    beego.Controller
}

func (c *AuthIndexController) Get() {
    sess := c.StartSession()
    
    username := sess.Get("username")
    userid := sess.Get("userid")
    
    logs.Debug(username, userid)
    
    c.Data["username"] = username
    c.Data["userid"] = userid
    logs.Debug("IP=>", c.Ctx.Input.IP())
    
    o := orm.NewOrm()
    user := models.User{Id: 1}
    
    err := o.Read(&user)
    
    if err == orm.ErrNoRows {
        logs.Debug("查询不到")
    } else if err == orm.ErrMissPK {
        logs.Debug("找不到主键")
    } else {
        logs.Debug(user.Id, user.Name)
    }
    
    c.Data["user"] = user
    logs.Debug(user)
    
    var users []*models.User
    num, err := o.QueryTable("go_user").All(&users)
    fmt.Printf("Returned Rows Num: %s, %s", num, err)
    c.Data["users"] = users
    
    c.TplName = "authIndex.html"
}


// 用户退出
type LogoutController struct {
    beego.Controller
}

func (c *LogoutController) Get() {
    sess := c.StartSession()
    //sess.Delete("userid")
    //sess.Delete("username")
    sess.Flush()
    c.Redirect("/login", 302)
}