package controllers

import (
	"Kate/features"
	models "Kate/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"os"
	"strconv"
)

type RegController struct {
	beego.Controller
}

func (this *RegController) Get() {

	sess_username, _ := this.GetSession("username").(string)
	sess_userlang, _ := this.GetSession("userlang").(string)

	if sess_username == "" {
		this.Data["LogStr"], this.Data["LogURL"] = features.Strings("", sess_userlang)
		this.Data["Main"] = features.Translate("Главная", sess_userlang)
		this.Data["Lang"] = features.Translate("Язык", sess_userlang)
		this.Data["Login"] = features.Translate("Логин", sess_userlang)
		this.Data["Password"] = features.Translate("Пароль", sess_userlang)
		this.Data["Reg"] = features.Translate("Регистрация", sess_userlang)
		this.TplNames = "index.tpl"
	} else {
		this.Redirect("/", 302)
	}
}

func (this *RegController) Post() {
	o := orm.NewOrm()
	o.Using("default")
	user := models.User{}
	this.ParseForm(&user)
	valid := validation.Validation{}
	isValid, _ := valid.Valid(user)
	var users []*models.User
	o.QueryTable("users").Filter("name", user.Name).All(&users)
	if !isValid {
		this.Redirect("/", 302)
	} else {
		if len(users) == 0 {
			user.Password = features.Encrypt_hash(user.Password, nil)
			id, _ := o.Insert(&user)
			os.Mkdir(fmt.Sprintf("../Downloads/%d", user.Id), 466)
			this.SetSession("userid", user.Id)
			this.SetSession("username", user.Name)
			beego.Info(user.Id, user.Name)
			url := "/user/" + strconv.FormatInt(id, 10)
			this.Redirect(url, 302)
		} else {
			this.Redirect("/", 302)
		}
	}
}
