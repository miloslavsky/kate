package models

type User struct {
	Id       int64  `form:"-"`
	Name     string `form:"name,text,name:" valid:"MinSize(3);MaxSize(40)"`
	Password string `form:"password,text,password" valid:"MinSize(5), MaxSize(100)"`
}

func (a *User) TableName() string {
	return "users"
}
