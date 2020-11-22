package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id			int64
	Username	string
	Password	string
	CreateTime	time.Time //Create time
    UpdateTime	time.Time  //Update time
}

func init() {
	orm.RegisterModel(new(User))

	orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("default", "mysql", "root:@/ops_go?charset=utf8")
}

// Create dan Update
func SaveUserInfoById(m *User) (err error) {
	o := orm.NewOrm()
	o.Using("default")
    var num int64
    if m.Id == 0 {
        m.CreateTime = time.Now()
        m.UpdateTime = time.Now()
        if num, err = o.Insert(m); err == nil {
            logs.Info("Number of records insert in database:", num)
        }
    } else {
        var tmp *User
        tmp, err = GetUserInfoById(m.Id)

        if err == nil {

            //Modify the names of several parameters.
            tmp.Username = m.Username
            tmp.UpdateTime = time.Now()

            if num, err = o.Update(tmp); err == nil {
                logs.Info("Number of records updated in database:", num)
            }
        }
    }
    return
}

// Delete
func DeleteUserInfo(id int64) (err error) {
	o := orm.NewOrm()
	o.Using("default")
    v := User{Id: id}
    if err = o.Read(&v, "Id"); err == nil {
        if num, err := o.Delete(&User{Id: id}); err == nil {
            logs.Info("Number of records deleted in database:", num)
        }
    }
    return
}

// Query by id
func GetUserInfoById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	o.Using("default")
    v = &User{Id: id}
    if err = o.Read(v, "Id"); err == nil {
        return v, nil
    }
    return nil, err
}

//Query data
func QueryAllUserInfo() (dataList []interface{}, err error) {
    var list []User
	o := orm.NewOrm()
	o.Using("default")
    qs := o.QueryTable(new(User))
    //Inquire
    //Query data
    if _, err = qs.All(&list); err == nil {
        for _, v := range list {
            dataList = append(dataList, v)
        }
        return dataList, nil
    }
    return nil, err
}