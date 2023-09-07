package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/menghua6/time-store/api"
	"github.com/menghua6/time-store/db"
)

var port string

func main() {
	flag.StringVar(&db.MysqlUserName, "mysql-username", "", "mysql username")
	flag.StringVar(&db.MysqlPassword, "mysql-password", "", "mysql password")
	flag.StringVar(&db.MysqlHost, "mysql-host", "", "mysql host")
	flag.StringVar(&db.MysqlPort, "mysql-port", "", "mysql port")
	flag.StringVar(&db.MysqlDbName, "mysql-db-name", "", "mysql-db-name")
	flag.StringVar(&api.FeishuUrl,"feishu-url","","feishu url")
	flag.StringVar(&port, "port", "", "port")
	flag.Parse()
	
	// if err := db.InitMysql(); err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		response, err := api.Schedule()
		if err != nil {
			log.Println(err.Error())
			context.String(400, "请求错误")
		} else {
			context.String(200, response)
		}
	})

	r.GET("/reserve/:describe", func(context *gin.Context) {
		describe := context.Param("describe")
		response, err := api.Reserve(describe)
		if err != nil {
			log.Println(err.Error())
			context.String(400, "请求错误")
		} else {
			context.String(200, response)
		}
	})

	r.Run(port)
}
