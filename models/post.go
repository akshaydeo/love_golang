package models

import (
	"love_golang/dbHelper"
	"time"
)

type Post struct {
	Id   int64
	Post string

	PostedBy User

	TotalUpvotes int64
	UpvotedBy    []User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (post *Post) Save() error {
	return dbHelper.DB().Where(Post{Id: post.Id}).Assign(post).FirstOrCreate(&post).Error
}
