package routers

import (
	"github.com/MicroCaas/bee-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
