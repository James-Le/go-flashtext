package main

import (
	_ "flash-text-api/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:test@tcp(127.0.0.1:3306)/testdb?charset=utf8", 30)
	orm.DefaultTimeLoc = time.UTC
}

func main() {
	beego.Run()
}
