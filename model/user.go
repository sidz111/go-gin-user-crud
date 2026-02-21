package model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
