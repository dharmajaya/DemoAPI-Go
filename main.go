package main

import (
	"github.com/astaxie/beego"
        "github.com/astaxie/beego/orm"
        _ "github.com/astaxie/beego/session/redis"
        _ "github.com/go-sql-driver/mysql"
	_ "./routers"
	_ "./models"
       "./controllers"
)


var (
  db_write_Force, _ = beego.AppConfig.Bool("db_write_Force")
  db_write_LogVerbose, _ = beego.AppConfig.Bool("db_write_LogVerbose")
  db_write_ConnString string = beego.AppConfig.String("db_write_ConnString")
  db_write_DriverName string = beego.AppConfig.String("db_write_DriverName")
  db_write_MaxIdle, _ = beego.AppConfig.Int("db_write_MaxIdle")
  db_write_MaxConn ,_  = beego.AppConfig.Int("db_write_MaxConn")

  db_read_LogVerbose, _ = beego.AppConfig.Bool("db_read_LogVerbose")
  db_read_ConnString string = beego.AppConfig.String("db_read_ConnString")
  db_read_DriverName string = beego.AppConfig.String("db_read_DriverName")
  db_read_MaxIdle, db_read_MaxIdleErr = beego.AppConfig.Int("db_read_MaxIdle")
  db_read_MaxConn, db_read_MaxConnErr = beego.AppConfig.Int("db_read_MaxConn")
)


func init() {

  orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("default", db_read_DriverName, db_read_ConnString, db_read_MaxIdle, db_read_MaxConn)
  orm.RegisterDataBase("write", db_write_DriverName, db_write_ConnString, db_write_MaxIdle, db_write_MaxConn)


  // Drop table and re-create.
  force := db_write_Force

  // Print log.
  verbose := db_write_LogVerbose

  // Database alias, Drop table and re-create, Print log
  err := orm.RunSyncdb("write", force, verbose)
  if err != nil {
    beego.Error(err)
  }
}


func main() {
  beego.ErrorController(&controllers.ErrorController{})
  beego.Run()
}
