package notes

import (
	"github.com/go-playground/validator/v10"
	"github.com/w33h/Productivity-Tracker-API/business/notes/spec"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	"github.com/w33h/Productivity-Tracker-API/exception"
)

type RepositoryNotes interface {
	InsertNote(notes Notes) (id string, err error)
	UpdateNote(notes Notes) (err error)
	DeleteNote(id string) (err error)
	FindNotesByStatus(status string, userId string) (notes []Notes, err error)
	FindNotesById(id string) (notes *Notes, err error)
	FindNotesByTags(tags string, userId string) (notes []Notes, err error)
	FindAllNotes(userId string) (notes []Notes, err error)
}

type ServiceNotes interface {
	CreateNote(specNotes spec.UpsertNotesSpec, userId string) (id string, err error)
	UpdateNote(specNotes spec.UpsertNotesSpec, id string) (err error)
	DeleteNote(id string) (err error)
	GetNotesByStatus(status string, userId string) (notes []Notes, err error)
	GetNotesById(id string) (notes *Notes, err error)
	GetNotesByTags(tags string, userId string) (notes []Notes, err error)
	GetAllNotes(userId string) (notes []Notes, err error)
}

type serviceNotes struct {
	notesRepo RepositoryNotes
	userRepo  user.RepositoryUser
	validate  *validator.Validate
}

func NewNoteService(notesRepo RepositoryNotes, userRepo user.RepositoryUser) ServiceNotes {
	return &serviceNotes{
		notesRepo: notesRepo,
		userRepo:  userRepo,
		validate:  validator.New(),
	}
}

func (s *serviceNotes) CreateNote(specNotes spec.UpsertNotesSpec, userId string) (id string, err error) {
	specNotes.UserId = userId
	err = s.validate.Struct(specNotes)
	if err != nil {
		return id, exception.ErrInvalidSpec
	}

	_, err = s.userRepo.FindById(userId)
	if err != nil {
		return id, exception.ErrNotFound
	}

	newNotes := NewNotes(specNotes.Status, specNotes.Content, specNotes.Tags, userId)

	id, err = s.notesRepo.InsertNote(newNotes)
	if err != nil {
		return id, exception.ErrInternalServer
	}

	return id, nil
}

func (s *serviceNotes) UpdateNote(specNotes spec.UpsertNotesSpec, id string) (err error) {
	err = s.validate.Struct(specNotes)
	if err != nil {
		return err
	}

	oldNotes, err := s.notesRepo.FindNotesById(id)
	if err != nil {
		return exception.ErrNotFound
	}

	_, err = s.userRepo.FindById(specNotes.UserId)
	if err != nil {
		return exception.ErrNotFound
	}

	err = s.CheckAuthorization(oldNotes.UserId, specNotes.UserId)
	if err != nil {
		return exception.ErrUnauthorized
	}

	newNotes := oldNotes.ModifyNotes(specNotes.Status, specNotes.Content, specNotes.Tags)

	err = s.notesRepo.UpdateNote(newNotes)
	if err != nil {
		return exception.ErrInternalServer
	}

	return nil
}

func (s *serviceNotes) DeleteNote(id string) (err error) {
	err = s.notesRepo.DeleteNote(id)
	if err != nil {
		return exception.ErrInternalServer
	}

	return nil
}

func (s *serviceNotes) GetNotesByStatus(status string, userId string) (notes []Notes, err error) {
	notes, err = s.notesRepo.FindNotesByStatus(status, userId)
	if err != nil {
		return nil, exception.ErrInternalServer
	}

	return notes, nil
}

func (s *serviceNotes) GetNotesById(id string) (notes *Notes, err error) {
	notes, err = s.notesRepo.FindNotesById(id)
	if err != nil {
		return nil, exception.ErrInternalServer
	}

	return notes, nil
}

func (s *serviceNotes) GetNotesByTags(tags string, userId string) (notes []Notes, err error) {
	notes, err = s.notesRepo.FindNotesByTags(tags, userId)
	if err != nil {
		return nil, exception.ErrInternalServer
	}

	return notes, nil
}

func (s *serviceNotes) GetAllNotes(userId string) (notes []Notes, err error) {
	notes, err = s.notesRepo.FindAllNotes(userId)
	if err != nil {
		return nil, exception.ErrInternalServer
	}

	return notes, nil
}

func (s *serviceNotes) CheckAuthorization(userId, id string) (err error) {
	if userId != id {
		return exception.ErrUnauthorized
	}

	return nil
}
