package main

import (
	models "Kate/models"
	_ "Kate/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	dbname := beego.AppConfig.String("mysqldb")
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", user, passwd, dbname))
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Blogpost))
	orm.RegisterModel(new(models.Follower))
	orm.RegisterModel(new(models.Comment))
	/*err := orm.RunSyncdb("default", true, true)
	if err != nil {

		fmt.Println(err)
	}*/
}

func main() {
	beego.SessionOn = true
	beego.SessionName = "kate"
	beego.AddFuncMap("pos", positive)
	beego.Run()
}

func positive(i int64) (o int64) {
	o = -i
	return
}
