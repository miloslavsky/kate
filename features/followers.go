package features

import (
	models "Kate/models"
	"github.com/astaxie/beego/orm"
)

func UserFollowers(follow int64, lang string) (exist bool, ufollowers map[int64]string) {
	o := orm.NewOrm()
	o.Using("default")
	ufollowers = make(map[int64]string)
	var followers []*models.Follower
	var users []*models.User
	o.QueryTable("followers").Filter("follow", follow).All(&followers)
	exist = true
	for _, follower := range followers {
		o.QueryTable("users").Filter("id", follower.Follower).All(&users)
		ufollowers[follower.Follower] = users[0].Name
	}
	if len(ufollowers) == 0 {
		exist = false
		ufollowers[int64(0)] = Translate("Нет подписчиков", lang)
	}
	return
}

func UserFollowing(follower int64, lang string) (exist bool, ufollowing map[int64]string) {
	o := orm.NewOrm()
	o.Using("default")
	ufollowing = make(map[int64]string)
	var followers []*models.Follower
	var users []*models.User
	o.QueryTable("followers").Filter("follower", follower).All(&followers)
	exist = true
	for _, follow := range followers {
		o.QueryTable("users").Filter("id", follow.Follow).All(&users)
		ufollowing[follow.Follow] = users[0].Name
	}
	if len(ufollowing) == 0 {
		exist = false
		ufollowing[int64(0)] = Translate("Нет подписок", lang)
	}
	return
}
