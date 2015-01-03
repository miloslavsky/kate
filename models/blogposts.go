package models

type Blogpost struct {
	Id      int64  `form:"-"`
	Name    string `form:"name,text,name:" valid:"MinSize(1);MaxSize(100)"`
	Content string `form:"content,text,content" valid:"MinSize(1), MaxSize(2000)"`
	Owner   int64  `form:"owner,text,owner`
}

func (a *Blogpost) TableName() string {
	return "blogposts"
}
