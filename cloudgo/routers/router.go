package routers

import (
	"cloudgo/controllers"
	"github.com/astaxie/beego"
)
//add route,use diffrent controllers to handle different route
func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/:id", &controllers.UserController{})
}
