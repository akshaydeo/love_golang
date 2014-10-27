package models

import (
	"log"
	"love_golang/dbWrapper"
)

type User struct {
	Id            int64
	TwitterHandle string
	TwitterId     string
	AuthKey       string
}

func (user *User) Save() error {
	return dbWrapper.DB().Where(User{TwitterHandle: user.TwitterHandle}).FirstOrCreate(user).Error
}

func GetUserById(id int64) (*User, error) {
	log.Println("Finding for id", id)
	var user User
	dbWrapper.DB().Model(User{}).Where("id = ?", id).Find(&user)
	log.Println("Found", user.TwitterHandle)
	return &user, nil
}
