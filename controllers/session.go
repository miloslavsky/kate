package controllers

import (
	"Kate/features"
	models "Kate/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type SessionController struct {
	beego.Controller
}

var (
	sess_username string
	sess_uid      int64
	sess_content  string
	Counter       map[string]string
)

func (this *SessionController) Get() {
	sess_username, _ := this.GetSession("username").(string)
	sess_userlang, _ := this.GetSession("userlang").(string)

	if sess_username == "" {

		this.Data["LogStr"], this.Data["LogURL"] = features.Strings("", sess_userlang)
		this.Data["Main"] = features.Translate("Главная", sess_userlang)
		this.Data["Lang"] = features.Translate("Язык", sess_userlang)
		this.Data["Login"] = features.Translate("Логин", sess_userlang)
		this.Data["Password"] = features.Translate("Пароль", sess_userlang)
		this.Data["Enter"] = features.Translate("Войти", sess_userlang)
		this.TplNames = "login.tpl"

	} else {
		this.Redirect("/", 302)
	}
}

func (this *SessionController) Post() {

	login := this.GetString("name")
	password := this.GetString("password")
	sess_userlang, _ := this.GetSession("userlang").(string)
	this.Data["LogStr"], this.Data["LogURL"] = features.Strings("", sess_userlang)
	this.Data["Main"] = features.Translate("Главная", sess_userlang)
	o := orm.NewOrm()
	o.Using("default")
	var users []*models.User
	o.QueryTable("users").Filter("name", login).All(&users)
	if len(users) == 0 {
		this.Redirect("/login", 302)
	} else {
		user := users[0]
		if features.Validate_hash(user.Password, password) {
			this.SetSession("userid", user.Id)
			this.SetSession("username", user.Name)
			url := "/user/" + strconv.FormatInt(user.Id, 10)
			this.Redirect(url, 302)
		} else {
			this.Redirect("/login", 302)
		}
	}
}

func (this *SessionController) Delete() {
	this.SetSession("userid", 0)
	this.SetSession("username", "")
	this.Redirect("/login", 302)
}
