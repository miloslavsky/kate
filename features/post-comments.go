package features

import (
	models "Kate/models"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func PostComments(post int64) map[int64]*models.Comment {
	o := orm.NewOrm()
	o.Using("default")
	var comments []*models.Comment
	o.QueryTable("comments").Filter("post", post).All(&comments)
	var res = make(map[int64]*models.Comment)
	for _, comment := range comments {
		comment.Id *= -1
		res[comment.Id] = comment
	}
	return res
}
