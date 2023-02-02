package model

import "hex-base/internal/constant"

type Todo struct {
	ID   uint    `gorm:"primaryKey" json:"id"`
	Form []*Form `gorm:"foreignKey:todo_id" json:"form"`
	Name string  `json:"name"`
	Task string  `json:"task"`
}

func (Todo) TableName() string {
	return constant.Todo.String()
}
