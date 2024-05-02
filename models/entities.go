package models

type Todo struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Descrition string `json:"description"`
	Done       bool   `json:"done"`
}
