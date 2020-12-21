package main

import (
	_ "BitcoinWeb/routers"
	"BitcoinWeb/utils"
	"github.com/astaxie/beego"
)

func main() {

	utils.PrepareJSON("getblock","00000000839a8e6886ab5951d76f411475428afc90947ee320161bbf18eb6048")
	//fmt.Println(utils.PrepareJSON("getblock","00000000839a8e6886ab5951d76f411475428afc90947ee320161bbf18eb6048"))


	beego.Run()
}

