package main

import (
	_ "BitcoinWeb/routers"
	"github.com/astaxie/beego"
)

func main() {
	fmt.println("123456789")
	beego.Run()
}

