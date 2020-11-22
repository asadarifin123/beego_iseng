package routers

import (
	"iseng/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.IndexController{})
    beego.Router("/postmortem", &controllers.PostmortemController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/dashboard", &controllers.DashboardController{})
    beego.Router("/dashboard/postmortem", &controllers.DashPostController{})
    beego.Router("/dashboard/pages", &controllers.DashPagesController{})
    beego.Router("/dashboard/web", &controllers.DashWebController{})
    beego.Router("/dashboard/ssl", &controllers.DashSSLController{})
    beego.Router("/dashboard/cve", &controllers.DashCVEController{})
    beego.Router("/dashboard/backup", &controllers.DashBackupController{})
    beego.Router("/dashboard/user", &controllers.DashUserController{})
    //beego.Router("/dashboard/user", &controllers.UserInfoController{})
    beego.Router("/pembuat", &controllers.PembuatController{})
}