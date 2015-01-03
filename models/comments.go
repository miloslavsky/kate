package models

type Comment struct {
	Id        int64  `form:"-"`
	Content   string `form:"content,text,content" valid:"MinSize(1), MaxSize(2000)"`
	OwnerName string `form:"name,text,name"`
	Owner     int64  `form:"owner,text,owner"`
	Post      int64  `form:"post,text,post"`
}

func (a *Comment) TableName() string {
	return "comments"
}
