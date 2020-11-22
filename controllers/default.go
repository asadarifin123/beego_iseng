package controllers

import (
	"iseng/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

type IndexController struct {
	beego.Controller
}

type PostmortemController struct {
	beego.Controller
}

type LoginController struct {
	beego.Controller
}

type DashboardController struct{
	beego.Controller
}

type DashPostController struct{
	beego.Controller
}

type DashPagesController struct{
	beego.Controller
}

type DashWebController struct{
	beego.Controller
}

type DashSSLController struct{
	beego.Controller
}

type DashCVEController struct{
	beego.Controller
}

type DashBackupController struct{
	beego.Controller
}

type DashUserController struct{
	beego.Controller
}

type PembuatController struct{
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *IndexController) Get() {
	c.Data["message"] = "HOME"
	c.TplName = "index.html"
}

func (c *PostmortemController) Get() {
	c.Data["message"] = "POSTMORTEM"
	c.TplName = "postmortem.html"
}

func (c *LoginController) Get() {
	c.Data["message"] = "LOGIN"
	c.TplName = "login.html"
}

func (c *DashboardController) Get() {
	c.Data["message"] = "DASHBOARD"
	c.TplName = "admin/dashboard.html"
}

func (c *DashPostController) Get() {
	c.Data["message"] = "Postmortem"
	c.TplName = "admin/dashpost.html"
}

func (c *DashPagesController) Get() {
	c.Data["message"] = "Pages"
	c.TplName = "admin/dashpages.html"
}

func (c *DashWebController) Get() {
	c.Data["message"] = "Web Monitoring"
	c.TplName = "admin/dashweb.html"
}

func (c *DashSSLController) Get() {
	c.Data["message"] = "SSL Checker"
	c.TplName = "admin/dashssl.html"
}

func (c *DashCVEController) Get() {
	c.Data["message"] = "CVE Data"
	c.TplName = "admin/dashcve.html"
}

func (c *DashBackupController) Get() {
	c.Data["message"] = "Web Backup"
	c.TplName = "admin/dashbackup.html"
}

func (c *DashUserController) Get() {
	c.Data["message"] = "User"
	c.TplName = "admin/dashuser.html"
}

func (c *PembuatController) Get() {
	c.Data["message"] = "Pembuat"
	c.TplName = "pembuat.html"
}

type UserInfoController struct {
	beego.Controller
}

// Modify

func (c *UserInfoController) Edit() {
	// Get id
    id, _ := c.GetInt64("Id", 0)
    userInfo, err := models.GetUserInfoById(id)
    if err == nil {
        c.Data["UserInfo"] = userInfo
    } else {
        tmpUserInfo := &models.User{}
        c.Data["UserInfo"] = tmpUserInfo
    }
    c.TplName = "userInfo/edit.html"
}

// Delete
//delete
func (c *UserInfoController) Delete() {
    // Get id
    id, _ := c.GetInt64("Id", 0)
    if err := models.DeleteUserInfo(id); err == nil {
        c.Data["json"] = "ok"
    } else {
        c.Data["json"] = "error"
    }
    c.ServeJSON()
}

//save
func (c *UserInfoController) Save() {
    // Automatically resolve the binding to the object
    userInfo := models.User{}
    if err := c.ParseForm(&userInfo); err == nil {
        if err := models.SaveUserInfoById(&userInfo); err == nil {
            c.Data["json"] = ""
        } else {
            c.Data["json"] = "error"
        }
    } else {
        c.Data["json"] = "error"
    }
    c.ServeJSON()
}

// Return all data
func (c *UserInfoController) List() {

    dataList, err := models.QueryAllUserInfo()
    if err == nil {
        c.Data["List"] = dataList
    }
    logs.Info("dataList :", dataList)
    c.TplName = "dashuser.html"

}