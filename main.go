package main

import (
<<<<<<< HEAD
	"BitcoinWeb/db_mysql"
	_ "BitcoinWeb/routers"
	"github.com/astaxie/beego"
)

func main() {

//	utils.PrepareJSON("getblock","00000000839a8e6886ab5951d76f411475428afc90947ee320161bbf18eb6048")
	//fmt.Println(utils.PrepareJSON("getblock","00000000839a8e6886ab5951d76f411475428afc90947ee320161bbf18eb6048"))

	//1.连接数据库
	db_mysql.ConnectDB()

	beego.Run()


	
}
=======
	"beegoText/db_mysql"
	_ "beegoText/routers"
)

func main() {
	//1.连接数据库,
	db_mysql.ConnectDB()

 }
>>>>>>> origin/hrx

