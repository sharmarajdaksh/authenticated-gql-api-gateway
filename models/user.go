package models

import "gorm.io/gorm"

// User data for authentication
type User struct {
	gorm.Model
	Email string
}

// GithubUser data for users who log in via Github
type GithubUser struct {
	gorm.Model
	GithubID  string `gorm:"primaryKey"`
	Name      string
	AvatarURL string
	UserID    string
	User      User `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

// BasicUser data for log in via email password combinations
type BasicUser struct {
	gorm.Model
	password string
	UserID   string
	User     User `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
