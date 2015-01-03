package controllers

import (
	models "Kate/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	sess_username, _ := this.GetSession("username").(string)
	if sess_username == "" {
		this.Redirect("/registration", 302)
	} else {
		o := orm.NewOrm()
		o.Using("default")
		var users []*models.User
		o.QueryTable("users").Filter("name", sess_username).All(&users)
		userId := users[0].Id
		url := "/user/" + strconv.FormatInt(userId, 10)
		this.Redirect(url, 302)
	}
}
