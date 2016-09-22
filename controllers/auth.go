package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego/orm"
    "beegoWork/models"
    "fmt"
    //"gopkg.in/pg.v4"
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
    //user := models.User{Id: 1}
    //
    //err := o.Read(&user)
    //
    //if err == orm.ErrNoRows {
    //    logs.Debug("查询不到")
    //} else if err == orm.ErrMissPK {
    //    logs.Debug("找不到主键")
    //} else {
    //    logs.Debug(user.Id, user.Name)
    //}
    //
    //c.Data["user"] = user
    //logs.Debug(user)

    //var posts []*models.Post
    //qs := o.QueryTable("go_post")
    //qs.RelatedSel("go_user")
    //num1, err1 := qs.Filter("User__Name__contains", "user").All(&posts)
    //fmt.Printf("Post Rows Num: %s, %s", num1, err1)

    user := &models.User{}
    o.QueryTable("go_user").Filter("Id", 1).RelatedSel().One(user)
    // 自动查询到 Profile
    fmt.Println(user.Profile)
    // 因为在 Profile 里定义了反向关系的 User，所以 Profile 里的 User 也是自动赋值过的，可以直接取用。
    fmt.Println(user.Profile.User)

    //var users []*models.User
    //num, err := o.QueryTable("go_user").All(&users)
    //fmt.Printf("Returned Rows Num: %s, %s", num, err)
    //logs.Debug("posts=>", posts)
    //c.Data["users"] = users
    //c.Data["posts"] = posts

    //user1 := &models.Users{
    //    Name:   "user1-beego-auth",
    //    Emails: []string{"user11@admin", "user12@admin"},
    //}
    //err := models.Db.Create(user1)
    //if err != nil {
    //    panic(err)
    //}
    //fmt.Println(user1)

    src1 := []string{"one@example.com", "two@example.com"}
    //var dst []string
    //_, err := db.QueryOne(pg.Scan(pg.Array(&dst)), `SELECT ?`, pg.Array(src))
    //if err != nil {
    //    panic(err)
    //}
    //fmt.Println(dst)

    zhuxh1 := &models.Zhuxh{
        Name:   "zhuxh1-beego-auth",
        Email: src1,
    }
    err := models.Db.Create(zhuxh1)
    if err != nil {
        panic(err)
    }
    fmt.Println(zhuxh1)

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