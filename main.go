package main

import (
	"beegoText/db_mysql"
	_ "beegoText/routers"
)

func main() {
	//1.连接数据库,
	db_mysql.ConnectDB()

 }

