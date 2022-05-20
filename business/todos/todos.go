package todos

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id        string
	UserId    string
	Status    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}

func NewTodos(
	status string,
	content string,
	userId string) Todo {

	return Todo{
		Id:        uuid.New().String(),
		UserId:    userId,
		Status:    status,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
}

func (old *Todo) ModifyTodos(newContent, status string) Todo {
	return Todo{
		Id:        old.Id,
		UserId:    old.UserId,
		Status:    status,
		Content:   newContent,
		CreatedAt: old.CreatedAt,
		UpdatedAt: time.Now(),
		Deleted:   old.Deleted,
	}
}
