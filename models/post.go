package models

import (
	"love_golang/dbWrapper"
	"time"
)

type Post struct {
	Id   int64
	Post string `sql:"type:text"`

	PostedBy       User
	PostedByUserId int64

	TotalUpvotes int64
	UpvotedBy    []User `gorm:"many2many:post_upvotes"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (post *Post) Save() error {
	return dbWrapper.DB().Save(post).Error
}

func GetAllPosts() (*[]Post, error) {
	var posts []Post
	db := dbWrapper.DB()
	if err := db.Order("updated_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}
	for index := 0; index < len(posts); index++ {
		user, _ := GetUserById(posts[index].PostedByUserId)
		posts[index].PostedBy = *user
	}
	return &posts, nil
}
