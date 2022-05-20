package request

import "github.com/w33h/Productivity-Tracker-API/business/todos/spec"

type CreateRequestTodo struct {
	UserId  string
	Content string `json:"content"`
	Status  string `json:"status" validate:"min=1,max=3"`
}

func (req *CreateRequestTodo) ToSpecTodo() *spec.UpsertTodosSpec {
	return &spec.UpsertTodosSpec{
		UserId: req.UserId,
		Content: req.Content,
		Status:  req.Status,
	}
}