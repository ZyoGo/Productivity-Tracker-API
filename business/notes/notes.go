package notes

import (
	"time"

	"github.com/google/uuid"
)

type Notes struct {
	Id        string
	UserId    string
	Status    string
	Content   string
	Tags      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}

func NewNotes(
	status string,
	content string,
	tags string,
	userId string,
) Notes {

	return Notes{
		Id:        uuid.New().String(),
		UserId:    userId,
		Status:    status,
		Content:   content,
		Tags:      tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
}

func (old *Notes) ModifyNotes(
	newStatus string,
	newContent string,
	newTags string,
) Notes {

	return Notes{
		Id:        old.Id,
		UserId:    old.UserId,
		Status:    newStatus,
		Content:   newContent,
		Tags:      newTags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
}
