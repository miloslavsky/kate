package controllers

import (
	"Kate/features"
	models "Kate/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type FollowController struct {
	beego.Controller
}

func (this *FollowController) Follow() {
	o := orm.NewOrm()
	o.Using("default")
	followId, _ := this.GetInt64(":id")
	sess_id := this.GetSession("userid")
	if sess_id == nil {
		this.Redirect("/", 302)
	} else {
		sess_uid = sess_id.(int64)
	}
	if sess_uid == followId {
		this.Redirect("/", 302)
	} else {
		var followers []*models.Follower
		follower := models.Follower{}
		o.QueryTable("followers").Filter("follower", sess_id).Filter("follow", followId).All(&followers)
		if len(followers) == 1 {
			follower = *followers[0]
			o.Delete(&follower)
		} else {
			follower.Follow = followId
			follower.Follower = sess_uid
			o.Insert(&follower)
		}
		url := "/user/" + strconv.FormatInt(followId, 10)
		this.Redirect(url, 302)
	}
}

func (this *FollowController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	followId, _ := this.GetInt64(":id")
	beego.Info("followId", followId)
	sess_username, _ := this.GetSession("username").(string)
	sess_userlang, _ := this.GetSession("userlang").(string)
	if sess_username != "" {
		this.Data["LogStr"], this.Data["LogURL"] = features.Strings(sess_username, sess_userlang)
		this.Data["FolEx"], this.Data["AllFollowers"] = features.UserFollowers(followId, sess_userlang)
		beego.Info(this.Data["FolEx"], this.Data["AllFollowers"])
		this.Data["Main"] = features.Translate("Главная", sess_userlang)
		this.Data["Lang"] = features.Translate("Язык", sess_userlang)
		this.Data["Followers"] = features.Translate("Подписчики", sess_userlang)
	} else {
		this.Redirect("/", 302)
	}
	this.TplNames = "followers.tpl"
}

func (this *FollowController) Following() {
	o := orm.NewOrm()
	o.Using("default")
	followId, _ := this.GetInt64(":id")
	beego.Info("followId", followId)
	sess_username, _ := this.GetSession("username").(string)
	sess_userlang, _ := this.GetSession("userlang").(string)
	if sess_username != "" {
		this.Data["LogStr"], this.Data["LogURL"] = features.Strings(sess_username, sess_userlang)
		this.Data["FolEx"], this.Data["AllFollowers"] = features.UserFollowing(followId, sess_userlang)
		beego.Info(this.Data["FolEx"], this.Data["AllFollowers"])
		this.Data["Main"] = features.Translate("Главная", sess_userlang)
		this.Data["Lang"] = features.Translate("Язык", sess_userlang)
		this.Data["Followers"] = features.Translate("Подписки", sess_userlang)
	} else {
		this.Redirect("/", 302)
	}
	this.TplNames = "followers.tpl"
}
