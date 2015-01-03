package controllers

import (
	"Kate/features"
	models "Kate/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	var users []*models.User
	userId, _ := this.GetInt64(":id")
	o.QueryTable("users").Filter("id", userId).All(&users)
	sess_username, _ := this.GetSession("username").(string)
	sess_uid, _ := this.GetSession("userid").(int64)
	sess_userlang, _ := this.GetSession("userlang").(string)
	this.Data["Owner"] = false
	this.Data["Exist"] = true
	if len(users) > 0 {
		user := users[0]
		this.Data["Id"] = strconv.FormatInt(user.Id, 10)
		this.Data["Followers"] = features.Translate("Подписчики", sess_userlang)
		this.Data["Following"] = features.Translate("Подписки", sess_userlang)
		this.Data["Name"] = user.Name
		this.Data["UserPosts"] = features.UserPosts(user.Id)
		if user.Name == sess_username {
			beego.Info("owner")
			this.Data["Owner"] = true
			this.Data["Posting"] = features.Translate("Новый пост", sess_userlang)
		} else {
			beego.Info("no owner", userId, sess_uid)
			if sess_username != "" {
				var followers []*models.Follower
				o.QueryTable("followers").Filter("follower", sess_uid).Filter("follow", userId).All(&followers)
				if len(followers) == 1 {
					this.Data["Follow"] = features.Translate("Отписаться", sess_userlang)
				} else {
					this.Data["Follow"] = features.Translate("Подписаться", sess_userlang)
				}
			} else {
				//offline
				this.Data["Follow"] = ``
				this.Data["Exist"] = false
			}
		}
	} else {
		this.Data["Exist"] = false
		this.Data["Name"] = features.Translate("Пользователя не существует", sess_userlang)
		this.Data["Follow"] = ``
		this.Data["UserPosts"] = make(map[int64](map[string]string))
		this.Data["Followers"] = ``
	}
	this.Data["LogStr"], this.Data["LogURL"] = features.Strings(sess_username, sess_userlang)
	this.Data["Main"] = features.Translate("Главная", sess_userlang)
	this.Data["Lang"] = features.Translate("Язык", sess_userlang)
	this.TplNames = "user.tpl"
}
