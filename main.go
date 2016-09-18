package main

import (
	_ "beegoWork/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/logs"
    _ "github.com/astaxie/beego/session/redis"
    "beegoWork/filter"
)

//var globalSessions *session.Manager

// 注册ORM
func ormRegister() {
    // 设置默认DB
    err := orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@127.0.0.1:5432/beegodb?sslmode=disable")
    if err != nil {
        logs.Debug("RegisterDataBase Error=>", err)
    }
    orm.Debug = true
}

// 自动建表
func createTable() {
    name := "default"  // 数据库别名
    force := false  // 不强制建数据库
    verbose := true  // 打印建表过程
    err := orm.RunSyncdb(name, force, verbose)  // 建表
    if err != nil {
        beego.Error(err)
    }
}

// 日志配置
func logsConfig() {
    //logs := logs.NewLogger(10000)
    logs.SetLogger("console")
    logs.SetLogger("file", `{"filename":"log/debug.log"}`)
    logs.EnableFuncCallDepth(true)  // 调用的文件名和文件行号
    logs.Async()  // 设置异步输出
}

// Session配置
func sessionInit()  {
    //globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
    //go globalSessions.GC()
    beego.BConfig.WebConfig.Session.SessionProvider = "redis"
    // 链接地址，连接池，访问密码(没有保持为空)
    beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379,0"
}


func init() {
    // 注册orm
    ormRegister()

    // 自动建表
    createTable()

    // 日志配置
    logsConfig()

    // Session配置
    sessionInit()
}

func main() {
    // 是否登录过渡器
    beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterUser)

    // 运行程序2
    beego.Run()
}
