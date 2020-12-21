package controllers

import (
	"BitcoinWeb/utils"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {





	c.TplName="test_demo01.html"
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	//c.Ctx.WriteString(utils.PrepareJSON("getbestblockhash"))
	data:=utils.PrepareJSON("getblock","0000000024366925dd090c74139b0a50865c892685d31ab7fc75b166d9d269db")
	//service.UnmarshalBlock()

	//fmt.Println(data)



	c.Data["Data"]=data
	//fmt.Println(Data)

}

