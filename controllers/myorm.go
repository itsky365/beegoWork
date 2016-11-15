package controllers

import (
	"beegoWork/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MyOrmController struct {
	beego.Controller
}

func (c *MyOrmController) Get() {
	//o := orm.NewOrm()
	//o.Using("default") // 默认使用 default，你可以指定为其他数据库
	//
	//profile := new(models.Profile)
	//profile.Age = 33
	//
	//user := new(models.User)
	//user.Profile = profile
	//user.Name = "slene3"
	//
	//fmt.Println(o.Insert(profile))
	//fmt.Println(o.Insert(user))

	//user := new(models.User)
	//user.Name = "slene6"
	//
	//fmt.Println(o.Insert(user))
	//
	//user.Name = "Your6"
	//fmt.Println(o.Update(user))
	//fmt.Println(o.Read(user))
	//fmt.Println(o.Delete(user))

	//o := orm.NewOrm()
	//user := models.User{Id: 4}
	//
	//err := o.Read(&user)
	//
	//profile := models.Profile{Age:41}
	////profile.Age = 40
	//o.Insert(&profile)
	//user.Profile = &profile
	//o.Update(&user)
	//
	//if err == orm.ErrNoRows {
	//    fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//    fmt.Println("找不到主键")
	//} else {
	//    fmt.Println(user.Id, user.Name)
	//}

	//o := orm.NewOrm()
	//user := models.User{Name: "slene1"}
	//err := o.Read(&user, "Name")
	//fmt.Println(&user)
	//if err != nil {
	//    fmt.Println(err)
	//}

	//o := orm.NewOrm()
	//user := models.User{Name: "slene"}
	//// 三个返回参数依次为：是否新创建的，对象Id值，错误
	//if created, id, err := o.ReadOrCreate(&user, "Name"); err == nil {
	//    if created {
	//        fmt.Println("New Insert an object. Id:", id)
	//    } else {
	//        fmt.Println("Get an object. Id:", id)
	//    }
	//}

	//o := orm.NewOrm()
	//var user models.User
	//user.Name = "sleneA"
	//
	//id, err := o.Insert(&user)
	//if err == nil {
	//    fmt.Println(id)
	//}

	o := orm.NewOrm()
	//user := models.User{Id: 1}
	//if o.Read(&user) == nil {
	//    user.Name = "MyName"
	//    if num, err := o.Update(&user, "Name"); err == nil {
	//        fmt.Println(num)
	//    }
	//}

	//user := &models.User{}
	////o.QueryTable("go_user").Filter("Id", 1).RelatedSel().One(user)
	//qs := o.QueryTable("go_user")
	////qs.Filter("id", 1) // WHERE id = 1
	////qs.Filter("profile__age__in", 31, 32, 33, 40)
	//qs.Filter("Id", 1).RelatedSel().All(user)
	//fmt.Println(user)
	//fmt.Println("name, age=>", user.Name, user.Profile, user.Profile.Age)

	//user := new(models.User)
	//user.Id = 2
	//o.Read(user)
	//fmt.Println(user)
	//
	//post := new(models.Post)
	//post.Title = "我是post2"
	//post.User = user
	//o.Insert(post)

	var post []*models.Post
	qs := o.QueryTable("go_post")
	qs.Filter("user__id__in", 1, 2, 3, 4).RelatedSel().All(&post)
	fmt.Println(post)
	c.Data["posts"] = post

	// my_news insert
	var my_user models.MyNews
	my_user.Title = "我的新闻标题"
	tags := map[string]string{
		"aa": "我很好",
		"bb": "我真的很棒",
	}
	tags_json, err := json.Marshal(tags)
	fmt.Println(tags_json)
	fmt.Println(string(tags_json))
	if err != nil {
		fmt.Println(err)
	}

	my_user.Tags = string(tags_json)
	my_id, err := o.Insert(&my_user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("my_id=>", my_id)

	var my_users []*models.MyNews
	num, err := o.QueryTable("go_my_news").All(&my_users)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	c.Data["my_users"] = my_users

	//c.Ctx.WriteString("hello")
	c.TplName = "myorm.html"
}
