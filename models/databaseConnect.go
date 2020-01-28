/**
 * @Author: liubaoshuai3
 * @Date: 2020/1/28 13:24
 * @File: databaseConnect
 * @Description:
 */

package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	dbuser := beego.AppConfig.String("mysqluser")
	dbpass := beego.AppConfig.String("mysqlpass")
	dbhost := beego.AppConfig.String("mysqlhost")
	dbport := beego.AppConfig.String("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	dataSource := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("read", "mysql", dataSource, 30)
}