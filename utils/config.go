package utils

import (
	"os"

	"github.com/astaxie/beego/config"
)

var (
	//G_server_addr  string //服务器ip地址
	G_server_port  string //服务器端口
	G_redis_addr   string //redis ip地址
	G_redis_port   string //redis port端口
	G_redis_dbnum  string //redis db 编号
	G_mysql_addr   string //mysql ip 地址
	G_mysql_port   string //mysql 端口
	G_mysql_dbname string //mysql db name
	G_image_addr   string
)

func InitConfig() {
	//从配置文件读取配置信息
	env := os.Getenv("ENV_CLUSTER")
	ConfPath := ""
	if env == "online" {
		ConfPath = "./conf/online.conf"
	} else if env == "beta" {
		ConfPath = "./conf/beta.conf"
	} else {
		ConfPath = "./conf/dev.conf"
	}
	appconf, _ := config.NewConfig("ini", ConfPath)
	// if err != nil {
	// 	beego.Debug(err)
	// 	return
	// }
	G_image_addr = appconf.String("imageaddr")
	//G_server_addr = appconf.String("httpaddr")
	G_server_port = appconf.String("httpport")
	G_redis_addr = appconf.String("redisaddr")
	G_redis_port = appconf.String("redisport")
	G_redis_dbnum = appconf.String("redisdbnum")
	G_mysql_addr = appconf.String("mysqladdr")
	G_mysql_port = appconf.String("mysqlport")
	G_mysql_dbname = appconf.String("mysqldbname")

	return
}

func init() {
	InitConfig()
}
