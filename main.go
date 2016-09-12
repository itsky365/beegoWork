package main

import (
    //"fmt"
    _ "beegoWork/routers"
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq"
    "github.com/astaxie/beego/logs"
    "github.com/astaxie/beego"
    //"beegoWork/models"
)

// Model Struct


func init() {
    // set default database
    err := orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@127.0.0.1:5432/canadaoffer?sslmode=disable")
    if err != nil {
        logs.Debug("RegisterDataBase error ", err)
    }

    // register model
    //orm.RegisterModel(new(models.User))

    orm.Debug = true

    // create table
    orm.RunSyncdb("default", false, true)
}

func main() {
    // delete
    //num, err = o.Delete(&u)
    //fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    beego.Run()
}
