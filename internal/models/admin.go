package models

type Admin struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
