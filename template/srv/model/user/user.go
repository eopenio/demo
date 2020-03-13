package model

import orm "github.com/eopenio/borm"

type User struct {
	UserId   int64  `json:"user_id" model:"user_id"`
	UserName string `json:"user_name" model:"user_name"`
	Password string `json:"password" model:"password"`
	CreateAt string `json:"create_at" model:"create_at"`
	Email    string `json:"email" model:"email"`
	Phone    string `json:"phone" model:"phone"`
}

func (user *User) TableName() string {
	return "user"
}

func init()  {
	orm.RegisterModel(new(User))
}

func (user *User) GetUser() (err error) {
	orm.NewOrm()
}