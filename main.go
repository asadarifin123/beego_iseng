package main

import (
    _ "iseng/routers"
    //"github.com/astaxie/beego/orm"
    //"fmt"
    "github.com/astaxie/beego"
    //_ "github.com/go-sql-driver/mysql"
)

// func connectdb()  {
// 	o := orm.NewOrm()
//     o.Using("default") // Using default, you can use other database

//     profile := new(models.Profile)
//     profile.Age = 30

//     user := new(models.User)
//     user.Profile = profile
//     user.Name = "slene"

//     fmt.Println(o.Insert(profile))
//     fmt.Println(o.Insert(user))
// }

func main() {
	beego.Run()
}

