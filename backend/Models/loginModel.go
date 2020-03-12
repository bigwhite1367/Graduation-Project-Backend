package models

type Login struct {
	UserID string `form:"userID" json:"userID" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	UserID     string
	UserName   string
	Privilege  string
}