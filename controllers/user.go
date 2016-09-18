package controllers

import (
    "github.com/astaxie/beego"
    "beegoWork/models"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/logs"
    "strconv"
)

type UserAddController struct {
    beego.Controller
}

func (c *UserAddController) Get() {
    c.TplName = "userAdd.html"
}

func (c *UserAddController) Post() {
    name := c.GetString("name")
    
    o := orm.NewOrm()
    
    var user models.User
    user.Name = name
    
    id, err := o.Insert(&user)
    if err == nil {
        logs.Debug(id)
    }
    c.Redirect("/auth/index", 302)
}


// 编辑用户
type UserEditController struct {
    beego.Controller
}

func (c *UserEditController) Get() {
    id := c.Ctx.Input.Param(":id")
    logs.Debug("param ", id)
    
    id2, err2 := strconv.Atoi(id)
    if err2 != nil {
        logs.Debug("edit atoi err=>", err2)
    }
    logs.Debug("id2=>", id2)
    
    o := orm.NewOrm()
    user := models.User{Id: id2}

    err := o.Read(&user)

    if err == orm.ErrNoRows {
        logs.Debug("查询不到")
    } else if err == orm.ErrMissPK {
        logs.Debug("找不到主键")
    } else {
        logs.Debug(user.Id, user.Name)
    }
    
    c.Data["user"] = user
    
    c.TplName = "userEdit.html"
}

func (c *UserEditController) Post() {
    logs.Debug("user edit input=>", c.Ctx.Input)
    
    id := c.Ctx.Input.Param(":id")
    logs.Debug("param ", id)
    name := c.GetString("name")
    
    id2, err2 := strconv.Atoi(id)
    if err2 != nil {
        logs.Debug("edit atoi err=>", err2)
    }
    logs.Debug("id2=>", id2)
    
    o := orm.NewOrm()
    user := models.User{Id: id2}
    
    err := o.Read(&user)
    
    if err == orm.ErrNoRows {
        logs.Debug("查询不到")
    } else if err == orm.ErrMissPK {
        logs.Debug("找不到主键")
    } else {
        logs.Debug(user.Id, user.Name)
    }
    user.Name = name
    
    r, err := o.Update(&user)
    if err != nil {
        logs.Debug("user edit err=>", err)
    }
    logs.Debug(r)
    
    c.Redirect("/auth/index", 302)
}