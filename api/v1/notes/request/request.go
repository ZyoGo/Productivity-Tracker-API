package request

import "github.com/w33h/Productivity-Tracker-API/business/notes/spec"

type CreateRequestNotes struct {
	Status  string `json:"status" validate:"required,alpha"`
	Content string `json:"content" validate:"required"`
	Tags    string `json:"tags"`
}

func (req *CreateRequestNotes) ToSpecNotes() *spec.UpsertNotesSpec {
	return &spec.UpsertNotesSpec{
		Status:  req.Status,
		Content: req.Content,
		Tags:    req.Tags,
	}
}
