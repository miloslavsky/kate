package features

import (
	models "Kate/models"
	"github.com/astaxie/beego/orm"
)

func UserPosts(owner int64) map[int64](map[string]string) {
	o := orm.NewOrm()
	o.Using("default")
	var userposts = make(map[int64](map[string]string))
	var posts []*models.Blogpost
	o.QueryTable("blogposts").Filter("owner", owner).All(&posts)
	for _, post := range posts {
		var userpost = make(map[string]string)
		userpost[post.Name] = post.Content
		userposts[-post.Id] = userpost
	}
	return userposts
}
