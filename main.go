package main

import (
	_ "ihome/routers"
	"net/http"
	"strings"

	_ "ihome/models"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	_ "github.com/beego/beego/v2/server/web/session/redis"
)

// func init() {
// 	orm.RegisterDriver("mysql", orm.DRMySQL)

// 	// set default database
// 	err := orm.RegisterDataBase("default", "mysql", "root:mysql@tcp("+utils.G_mysql_addr+":"+utils.G_mysql_port+")/"+utils.G_mysql_dbname+"?charset=utf8")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// fmt.Println("innn")
// 	// create table
// 	//第二个参数是强制更新数据库
// 	//第三个参数是如果没有则同步
// 	// orm.RunSyncdb("default", false, true)
// }

func main() {

	ignoreStaticPath()
	beego.Run()

}

func ignoreStaticPath() {
	beego.InsertFilter("/", beego.BeforeRouter, TransportentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransportentStatic)
	beego.InsertFilter("/api/*", beego.AfterExec, setSession)
	beego.InsertFilter("/api/v1.0/user/*", beego.BeforeExec, checkLogin)
	beego.InsertFilter("/api/v1.0/user", beego.BeforeExec, checkLogin)
}

func TransportentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
func setSession(ctx *context.Context) {
	id, ok := ctx.Input.Session("id").(int)
	if ok {
		ctx.Output.Session("id", id)
	}
}
func checkLogin(ctx *context.Context) {
	// fmt.Println(ctx.Request.Method, ctx.Request.URL.Path)
	if ctx.Request.Method == "POST" && ctx.Request.URL.Path == "/api/v1.0/user" {
		return
	}
	_, ok := ctx.Input.Session("id").(int)
	if !ok {
		resp := make(map[string]interface{})
		resp["errno"] = 4101
		resp["errmsg"] = "请先登录"
		ctx.Output.JSON(resp, false, false)

	}
}
