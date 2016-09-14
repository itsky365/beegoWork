package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id      int    `orm:"pk;auto"`
	Name    string `orm:"size(100)"`
	Created time.Time
	//Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	//Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

//定义结构体，名字为表名大写，字段大写，为表的字段
type Book struct {
	Name   string
	Num    int64 `orm:"pk;auto"` //主键，自动增长
	Author string
	Price  float32
}

//type Profile struct {
//	Id   int
//	Age  int16
//	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
//}
//
//type Post struct {
//	Id    int
//	Title string
//	User  *User  `orm:"rel(fk)"` //设置一对多关系
//	Tags  []*Tag `orm:"rel(m2m)"`
//}
//
//type Tag struct {
//	Id    int
//	Name  string
//	Posts []*Post `orm:"reverse(many)"`
//}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix("go_", new(User)) //带前缀的表
	orm.RegisterModelWithPrefix("go_", new(Book))
	//orm.RegisterModel(new(User))
}
