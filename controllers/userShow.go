package controllers

import (
	"BitcoinWeb/entity"
	"BitcoinWeb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type ShowTime struct {
	beego.Controller
}





func (s *ShowTime) Post()  {
	var comd entity.RpcCommand
	err := s.ParseForm(&comd)
	if err != nil {
		s.Ctx.WriteString("输入命令错误！")
		return
	}
	fmt.Println("命令为：", comd.Command)
	fmt.Println("参数一为：", comd.First_parameter)
	fmt.Println("参数二为：",comd.Second_parameter)

	//if comd.First_parameter==0 {
	//	data1:=utils.PrepareJSON(comd.Command,comd.Second_parameter)
	//	t.Data["Data"]=data1
	//}else {
	//
	//	data:=utils.PrepareJSON(comd.Command,comd.First_parameter)
	//
	//	t.Data["Data"]=data
	//}

	var data interface{}
	switch comd.Command {
	case "getblockhash" ,"getblockfilter":
		height, err := strconv.Atoi(comd.First_parameter)
		if err != nil {
			s.Ctx.WriteString("参数类型有误，请重新输入！")
			return
		}
		data = utils.PrepareJSON(comd.Command, height)
	case "getblock", "gettransaction","addnode":
		data = utils.PrepareJSON(comd.Command, comd.First_parameter)
	case "getbestblockhash","getblockchaininfo","getblockcount","getchaintips","getdifficulty","getmempoolinfo","gettxoutsetinfo","savemempool","verifychain","getrpcinfo","help","stop","uptime","getmininginfo","getnewadress":
		data = utils.PrepareJSON(comd.Command)
	}




	//switch comd.Command {
	//case "":
	//	height, err := strconv.Atoi(comd.Second_parameter)
	//	if err != nil {
	//		s.Ctx.WriteString("参数类型有误，请重新输入！")
	//		return
	//	}
	//	data = utils.PrepareJSON(comd.Command, comd.First_parameter, height)
	//case "addnode":
	//	data = utils.PrepareJSON(comd.Command, comd.First_parameter, comd.Second_parameter)
	//
	//}

	s.Data["Data"] = data

	s.TplName = "show_time.html"
}
