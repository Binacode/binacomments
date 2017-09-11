package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserID   uint      `json: "userID"`
	ParentID uint      `json: "parentID"`
	Votes    int32     `json: "votes"`
	Content  string    `json: "content"`
	HasVote  int8      `json: "hasvote" gorm: "-"`
	User     []User    `json: "user,omitempty"`
	Children []Comment `json: "children,omitempty"`
}
