package main

import (
	"content_system/internal/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := connDB()
	fmt.Println(db)
}
func connDB() *gorm.DB {
	//user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	mysqlDB, er := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/cms_account?charset=utf8mb4&parseTime=True&loc=Local"))
	if er != nil {
		panic(er)
	}
	//拿到mysqlDB的实例
	mysqlDB.AutoMigrate(&model.Account{}) //自动迁移，自动创建表，默认蛇行负数
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100) //最大连接数
	db.SetMaxIdleConns(50)  //最大空闲连接数，一般为最大连接数/2
	//if env == "test" {
	//	mysqlDB = mysqlDB.Debug()
	//}
	mysqlDB = mysqlDB.Debug()
	return mysqlDB
}
