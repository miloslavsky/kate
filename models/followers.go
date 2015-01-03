package models

type Follower struct {
	Id       int64
	Follow   int64
	Follower int64
}

func (a *Follower) TableName() string {
	return "followers"
}
