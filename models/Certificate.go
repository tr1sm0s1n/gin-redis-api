package models

type Certificate struct {
	Id     string `json:"id" redis:"id" binding:"required"`
	Name   string `json:"name" redis:"name" binding:"required"`
	Course string `json:"course" redis:"course" binding:"required"`
	Grade  string `json:"grade" redis:"grade" binding:"required"`
	Date   string `json:"date" redis:"date" binding:"required"`
}
