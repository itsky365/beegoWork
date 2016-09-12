package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/astaxie/beego/orm"
	//"beegoWork/models"
	"beegoWork/models"
	"github.com/astaxie/beego/orm"
)

type BlogIndexController struct {
	beego.Controller
}

func (c *BlogIndexController) Get() {
	id := c.Ctx.Input.Param(":id")
	//fmt.Println(id)
	//fmt.Println(c.Ctx.Input.Params())

	page := c.GetString("page")
	//fmt.Println("page=>", page)
	logs.Debug("page", page)

	o := orm.NewOrm()

	user := models.User{Name: "slene22"}

	// insert
	user_id, err := o.Insert(&user)
	logs.Debug("ID: %d, ERR: %v\n", user_id, err)

	// update
	//user.Name = "astaxie22"
	//num, err := o.Update(&user)
	//logs.Debug("NUM: %d, ERR: %v\n", num, err)

	// read one
	//u := models.User{Id: user.Id}
	//err = o.Read(&u)
	//logs.Debug("ERR: %v\n", err)

	//o := orm.NewOrm()
	//logs.Debug(o)
	////o.Using("default") // 默认使用 default，你可以指定为其他数据库
	//user := models.User{Name: "我是谁2312"}
	//id, err := o.Insert(&user)
	//if err != nil {
	//    logs.Debug(err)
	//}
	//logs.Debug(err)

	c.Data["Id"] = id
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Layout = "layout.tpl"
	c.TplName = "blogIndex.tpl"
}

type BlogJsonController struct {
	beego.Controller
}

func (c *BlogJsonController) Get() {
	c.Data["json"] = map[string]string{
		"name": "爱wifi",
		"work": "chinanet",
	}
	c.ServeJSON()
}
