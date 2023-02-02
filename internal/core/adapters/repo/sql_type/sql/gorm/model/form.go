package model

import "hex-base/internal/constant"

type Form struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Category  string `gorm:"column:category" json:"category""`
	IsDeleted string `gorm:"column:is_deleted" json:"is_deleted"`
	Status    string `gorm:"column:status" json:"status"`
	TodoID    uint   `gorm:"column:todo_id" json:"todo_id"`
	Title     string `gorm:"column:title" json:"title"`
}


func (Form) TableName() string {
	return constant.Form.String()
}