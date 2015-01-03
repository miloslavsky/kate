package controllers

import (
	//"Kate/features"
	models "Kate/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Post() {
	o := orm.NewOrm()
	o.Using("default")
	comment := models.Comment{}
	this.ParseForm(&comment)
	sess_uid, _ := this.GetSession("userid").(int64)
	sess_username, _ := this.GetSession("username").(string)
	postId, _ := this.GetInt64(":id")
	comment.Owner = sess_uid
	comment.Post = postId
	comment.OwnerName = sess_username
	valid := validation.Validation{}
	isValid, _ := valid.Valid(comment)
	if !isValid {
		beego.Info("invalid")
	} else {
		o.Insert(&comment)
	}
	this.Redirect(fmt.Sprintf("/posts/%d", postId), 302)
}

func (this *CommentController) Delete() {
	o := orm.NewOrm()
	o.Using("default")
	comment := models.Comment{}
	commentId, _ := this.GetInt64(":id")
	var comments []*models.Comment
	o.QueryTable("comments").Filter("id", commentId).All(&comments)
	if len(comments) == 0 {
		this.Redirect("/", 302)
	}
	comment = *comments[0]
	if sess_id := this.GetSession("userid"); comment.Owner == sess_id {
		o.Delete(&comment)
	}
	this.Redirect(fmt.Sprintf("/posts/%d", comment.Post), 302)
}
