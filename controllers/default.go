package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
<<<<<<< HEAD


	c.TplName = "login.html"



	//c.TplName="test_demo01.html"
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	//c.Ctx.WriteString(utils.PrepareJSON("getbestblockhash"))

	//service.UnmarshalBlock()


	//	data:=utils.PrepareJSON("getblock","0000000024366925dd090c74139b0a50865c892685d31ab7fc75b166d9d269db")
	//fmt.Println(data)

	//c.Data["Data"]=data
	//fmt.Println(Data)

}

=======
	c.TplName = "login.html"
}
>>>>>>> origin/hrx
