package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Password  string `json:"password" gorm:"not null"`
	Name      string `json:"name"`
	UrlAvatar string `json:"url_avatar"`
}
