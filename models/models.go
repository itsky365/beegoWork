package models

import (
	"github.com/astaxie/beego/orm"
	"time"
    //"github.com/lib/pq"
)

type User struct {
	Id      int    `orm:"pk;auto"`
	Name    string `orm:"size(100)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Profile *Profile `orm:"null;rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

//CREATE TABLE public.go_user (
//	id integer NOT NULL DEFAULT nextval('go_user_id_seq'::regclass),
//	name character varying(100) NOT NULL DEFAULT ''::character varying,
//	created timestamp with time zone NOT NULL,
//	profile_id integer,
//	CONSTRAINT go_user_pkey PRIMARY KEY (id),
//	CONSTRAINT go_user_profile_id_key UNIQUE (profile_id)
//)


type Profile struct {
    Id   int `orm:"pk;auto"`
    Age  int16
    User *User `orm:"reverse(one)"` //设置一对一反向关系(可选)
}

//CREATE TABLE public.go_profile (
//	id integer NOT NULL DEFAULT nextval('go_profile_id_seq'::regclass),
//	age smallint NOT NULL DEFAULT 0,
//	CONSTRAINT go_profile_pkey PRIMARY KEY (id)
//)


type Post struct {
    Id    int    `orm:"pk;auto"`
    Title string `orm:"size(100)"`
    User  *User  `orm:"rel(fk)"`
    Tags  []*Tag `orm:"rel(m2m)"`
}

//CREATE TABLE public.go_post (
//	id integer NOT NULL DEFAULT nextval('go_post_id_seq'::regclass),
//	title character varying(100) NOT NULL DEFAULT ''::character varying,
//	user_id integer NOT NULL,
//	CONSTRAINT go_post_pkey PRIMARY KEY (id)
//)

//CREATE TABLE public.go_post_go_tags (
//	id integer NOT NULL DEFAULT nextval('go_post_go_tags_id_seq'::regclass),
//	go_post_id integer NOT NULL,
//	go_tag_id integer NOT NULL,
//	CONSTRAINT go_post_go_tags_pkey PRIMARY KEY (id)
//)


type Tag struct {
    Id    int `orm:"pk;auto"`
    Name  string
    Posts []*Post `orm:"reverse(many)"`
}

//type Tag1 *pq.StringArray

type MyNews struct {
	Id    int `orm:"pk;auto"`
	Title  string
    Tags string `orm:"type(jsonb)"`
}

//CREATE TABLE public.go_tag
//(
//	id integer NOT NULL DEFAULT nextval('go_tag_id_seq'::regclass),
//	name text NOT NULL DEFAULT ''::text,
//	CONSTRAINT go_tag_pkey PRIMARY KEY (id)
//)


////定义结构体，名字为表名大写，字段大写，为表的字段
//type Book struct {
//	Name   string
//	Num    int64 `orm:"pk;auto"` //主键，自动增长
//	Author string
//	Price  float32
//}

// note
// `orm:"rel(pk)"` `orm:"rel(one)"` 生成对应的外键


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModelWithPrefix("go_", new(User)) //带前缀的表
    //orm.RegisterModel(new(User))
    //orm.RegisterModelWithPrefix("go_", new(Book))
	orm.RegisterModelWithPrefix("go_", new(Profile))
	orm.RegisterModelWithPrefix("go_", new(Post))
	orm.RegisterModelWithPrefix("go_", new(Tag))
	orm.RegisterModelWithPrefix("go_", new(MyNews))
}
