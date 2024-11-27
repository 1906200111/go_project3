package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
}

func main() {
	db := connDB()
}
func connDB() *gorm.DB {
	//user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	mysqlDB, er := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/cms_account?charset=utf8mb4&parseTime=True&loc=Local"))
	if er != nil {
		panic(er)
	}
	//拿到mysqlDB的实例
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
