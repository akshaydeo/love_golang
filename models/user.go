package models

import (
	"love_golang/dbHelper"
)

type User struct {
	Id            int64
	TwitterHandle string
	TwitterId     string
	AuthKey       string
}

func (user *User) Save() error {
	return dbHelper.DB().Save(&user).Error()
}
