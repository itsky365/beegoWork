package filter

import (
    "github.com/astaxie/beego/context"
)

//type FilterFunc func(*context.Context)

// 用户登录判断
var FilterUser = func(ctx *context.Context) {
    _, ok := ctx.Input.Session("userid").(int)
    if !ok && ctx.Request.RequestURI != "/login" {
        ctx.Redirect(302, "/login")
    }
}