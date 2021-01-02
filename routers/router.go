package routers

import (
<<<<<<< HEAD
	"BitcoinWeb/controllers"
=======
	"beegoText/controllers"
>>>>>>> origin/hrx
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
    beego.Router("/", &controllers.MainController{})
<<<<<<< HEAD
=======

>>>>>>> origin/hrx
	//用户注册的接口请求
	beego.Router("/user_register",&controllers.RegisterController{})

	//用户登录的页面请求接口
	beego.Router("/login.html",&controllers.LoginController{})

	//用户登录请求接口
	beego.Router("/user_login",&controllers.LoginController{})

<<<<<<< HEAD

//	beego.Router("/test",&controllers.Test{})


    beego.Router("/show_time",&controllers.ShowTime{})
=======
>>>>>>> origin/hrx
}
