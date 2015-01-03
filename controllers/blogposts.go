package controllers

import (
	"Kate/features"
	models "Kate/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type BlogPostController struct {
	beego.Controller
}

func (this *BlogPostController) Add() {
	sess_username, _ := this.GetSession("username").(string)
	sess_userlang, _ := this.GetSession("userlang").(string)
	if sess_username == "" {
		this.Redirect("/", 302)
	}
	this.Data["LogStr"], this.Data["LogURL"] = features.Strings(sess_username, sess_userlang)
	this.Data["Main"] = features.Translate("Главная", sess_userlang)
	this.Data["Lang"] = features.Translate("Язык", sess_userlang)
	this.Data["Topic"] = features.Translate("Тема", sess_userlang)
	this.Data["BlogPost"] = features.Translate("Пост", sess_userlang)
	this.Data["Content"] = features.Translate("Содержимое", sess_userlang)
	this.Data["Send"] = features.Translate("Отправить", sess_userlang)
	this.Data["New"] = features.Translate("Новый пост", sess_userlang)
	this.TplNames = "new-post.tpl"
}

func (this *BlogPostController) Post() {
	o := orm.NewOrm()
	o.Using("default")
	post := models.Blogpost{}
	this.ParseForm(&post)
	sess_id, _ := this.GetSession("userid").(int64)
	post.Owner = sess_id
	valid := validation.Validation{}
	isValid, _ := valid.Valid(post)
	if !isValid {
		this.Redirect("/add", 302)
	} else {
		o.Insert(&post)
		this.Redirect("/", 302)
	}
}

func (this *BlogPostController) Get() {
	sess_username, _ := this.GetSession("username").(string)
	sess_userlang, _ := this.GetSession("userlang").(string)
	sess_uid, _ := this.GetSession("userid").(int64)
	if sess_username == "" {
		this.Redirect("/", 302)
	}
	o := orm.NewOrm()
	o.Using("default")
	var posts []*models.Blogpost
	postId, _ := this.GetInt(":id")
	o.QueryTable("blogposts").Filter("id", postId).All(&posts)
	if len(posts) > 0 {
		post := posts[0]
		this.Data["PostName"] = post.Name
		this.Data["Content"] = post.Content
		this.Data["LogStr"], this.Data["LogURL"] = features.Strings(sess_username, sess_userlang)
		this.Data["Main"] = features.Translate("Главная", sess_userlang)
		this.Data["Lang"] = features.Translate("Язык", sess_userlang)
		this.Data["Comment"] = features.Translate("Комментарий", sess_userlang)
		this.Data["Send"] = features.Translate("Отправить", sess_userlang)
		this.Data["Delete"] = features.Translate("Удалить", sess_userlang)
		this.Data["Edit"] = features.Translate("Редактировать", sess_userlang)
		this.Data["PostComments"] = features.PostComments(int64(postId))
		this.Data["Sess"] = sess_uid
		beego.Info(this.Data["PostComments"])
		if sess_uid == post.Owner {
			this.Data["Owner"] = true
			this.Data["Id"] = strconv.Itoa(postId)
			this.Data["DeletePost"] = features.Translate("Удалить пост", sess_userlang)
		} else {
			this.Data["Owner"] = false
		}
		this.TplNames = "blogpost.tpl"
	} else {
		this.Redirect("/", 302)
	}
}

func (this *BlogPostController) Delete() {
	o := orm.NewOrm()
	o.Using("default")
	post := models.Blogpost{}
	postId, _ := this.GetInt64(":id")
	var posts []*models.Blogpost
	o.QueryTable("blogposts").Filter("id", postId).All(&posts)
	if len(posts) == 0 {
		this.Redirect("/", 302)
	}
	post = *posts[0]
	if sess_id := this.GetSession("userid"); post.Owner == sess_id {
		o.Delete(&post)
		var comments []*models.Comment
		o.QueryTable("comments").Filter("post", postId).All(&comments)
		for _, comment := range comments {
			o.Delete(&comment)
		}
	}
	this.Redirect("/", 302)
}

func (this *BlogPostController) Edit() {
	sess_username, _ := this.GetSession("username").(string)
	sess_userlang, _ := this.GetSession("userlang").(string)
	if sess_username == "" {
		this.Redirect("/", 302)
	}
	o := orm.NewOrm()
	o.Using("default")
	postId, _ := this.GetInt64(":id")
	var posts []*models.Blogpost
	o.QueryTable("blogposts").Filter("id", postId).All(&posts)
	if len(posts) == 0 {
		this.Redirect("/", 302)
	}
	if sess_id := this.GetSession("userid"); posts[0].Owner != sess_id {
		this.Redirect("/", 302)
	}
	this.Data["LogStr"], this.Data["LogURL"] = features.Strings(sess_username, sess_userlang)
	this.Data["Main"] = features.Translate("Главная", sess_userlang)
	this.Data["Lang"] = features.Translate("Язык", sess_userlang)
	this.Data["Topic"] = posts[0].Name
	this.Data["BlogPost"] = features.Translate("Пост", sess_userlang)
	this.Data["Content"] = posts[0].Content
	this.Data["Edit"] = features.Translate("Редактировать", sess_userlang)
	this.TplNames = "edit-post.tpl"
}

func (this *BlogPostController) Edition() {
	o := orm.NewOrm()
	o.Using("default")
	post := models.Blogpost{}
	postId, _ := this.GetInt64(":id")
	var posts []*models.Blogpost
	o.QueryTable("blogposts").Filter("id", postId).All(&posts)
	if len(posts) == 0 {
		this.Redirect("/", 302)
	}
	post = *posts[0]
	if sess_id := this.GetSession("userid"); post.Owner == sess_id {
		this.ParseForm(&post)
		valid := validation.Validation{}
		isValid, _ := valid.Valid(post)
		if !isValid {
			beego.Info("invalid")
			this.Redirect("/add", 302)
		} else {
			o.Update(&post)
			this.Redirect("/", 302)
		}
	} else {
		this.Redirect("/add", 302)
	}

}
