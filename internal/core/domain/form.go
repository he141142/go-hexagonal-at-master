package domain

import (
	"context"
	"hex-base/internal/constant"
)

type (
	Form struct {
		Id        uint
		Category  string
		IsDeleted bool
		Status    string
		TodoID    uint
		Title     string
		Todo      *Todo
	}

	FormList struct {
		Data []*Form
	}

	FormRepository interface {
		Create(context.Context, Form) (*Form, error)
		GetById(context.Context, uint) (*Form, error)
		ListByTodoId(context.Context, uint) (*FormList,error)
	}
)

func NewForm(id uint, category string, isDeleted bool, status string, todoID uint, title string) *Form {
	return &Form{
		Id:        id,
		Category:  category,
		IsDeleted: isDeleted,
		Status:    status,
		TodoID:    todoID,
		Title:     title,
	}
}

func (Form) TableName() string {
	return constant.Form.String()
}

func (f *Form) GetId() uint {
	return f.Id
}

func (f *Form) GetCategory() string {
	return f.Category
}

func (f *Form) GetStatus() string {
	return f.Status
}

func (f *Form) GetTodoId() uint {
	return f.TodoID
}

func (f *Form) GetTitle() string {
	return f.Title
}
