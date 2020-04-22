package routers

import (
	"flash-text-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/all", &controllers.TextController{}, "get:ListWords")
	beego.Router("/addAll", &controllers.TextController{}, "post:AddAllWords")
	beego.Router("/parse", &controllers.TextController{}, "post:Parse")
}
