package controllers

import (
	"beegoWork/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type BlogIndexController struct {
	beego.Controller
}

type A struct {
	Name string
	Age  int
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
	fmt.Println(user_id)
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

	//books := []models.Book{
	//	{Name: "lkn1", Price: 10},
	//	{Name: "lkn5", Price: 122},
	//	{Name: "lkn3", Price: 3},
	//	{Name: "lkn22", Price: 12},
	//	{Name: "lkn2121", Price: 133},
	//}
	//插入，第一个参数为要插入的条数
	//booksId, err := o.InsertMulti(5, books)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(booksId)

	c.Data["AA"] = &A{Name: "astaxie", Age: 25}

	ss :=[]string{"a","b","c"}
	c.Data["s"]=ss

	c.Data["Id"] = id
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["html"] = "<div style='color:#f00'>你好html</div>"

	c.Layout = "layout.html"
	c.TplName = "blogIndex.html"
}

type BlogJsonController struct {
	beego.Controller
}

func (c *BlogJsonController) Get() {
	o := orm.NewOrm()

	var book1 *models.Book = &models.Book{Name: "蓝色fasfsaf天空", Price: 30.0, Author: "刘凯宁"}
	//间接构造对象
	var book2 *models.Book = new(models.Book)
	book2.Author = "mq"
	book2.Name = "mq111fasf"
	book2.Price = 100.0
	//执行插入，o.Insert()返回的是受影响的id
	id1, _ := o.Insert(book1)
	id2, _ := o.Insert(book2)
	logs.Debug(id1)
	logs.Debug(id2)
	//logs.Debug("blogjson=>", o.Insert(book1))
	//logs.Debug("blogjson=>", o.Insert(book2))

	c.Data["json"] = map[string]string{
		"name": "爱wifi",
		"work": "chinanet",
	}
	c.ServeJSON()
}
