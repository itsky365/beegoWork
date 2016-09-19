package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id      int    `orm:"pk;auto"`
	Name    string `orm:"size(100)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Profile *Profile `orm:"null;rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
    Id   int `orm:"pk;auto"`
    Age  int16
    User *User `orm:"reverse(one)"` //设置一对一反向关系(可选)
}

type Post struct {
    Id    int    `orm:"pk;auto"`
    Title string `orm:"size(100)"`
    User  *User  `orm:"rel(fk)"`
    Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
    Id    int `orm:"pk;auto"`
    Name  string
    Posts []*Post `orm:"reverse(many)"`
}

////定义结构体，名字为表名大写，字段大写，为表的字段
//type Book struct {
//	Name   string
//	Num    int64 `orm:"pk;auto"` //主键，自动增长
//	Author string
//	Price  float32
//}


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix("go_", new(User)) //带前缀的表
    //orm.RegisterModel(new(User))
    //orm.RegisterModelWithPrefix("go_", new(Book))
	orm.RegisterModelWithPrefix("go_", new(Profile))
	orm.RegisterModelWithPrefix("go_", new(Post))
	orm.RegisterModelWithPrefix("go_", new(Tag))
}
