package notes

import (
	domain "github.com/w33h/Productivity-Tracker-API/business/notes"
	"gorm.io/gorm"
)

type notesRepository struct {
	db *gorm.DB
}

func NewNotesRepository(db *gorm.DB) domain.RepositoryNotes {
	return &notesRepository{db}
}

func (r *notesRepository) InsertNote(notes domain.Notes) (id string, err error) {
	if err = r.db.Create(&notes).Error; err != nil {
		return id, err
	}

	id = notes.Id

	return id, nil
}

func (r *notesRepository) UpdateNote(notes domain.Notes) (err error) {
	if err = r.db.Save(&notes).Error; err != nil {
		return err
	}

	return nil
}

func (r *notesRepository) DeleteNote(id string) (err error) {
	if err = r.db.Table("notes").Where("Id = ?", id).Update("deleted", true).Error; err != nil {
		return err
	}

	return nil
}

func (r *notesRepository) FindNotesByStatus(status string, userId string) (notes []domain.Notes, err error) {
	if err = r.db.Where("status = ? AND deleted = ? AND user_id = ?", status, false, userId).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *notesRepository) FindNotesById(id string) (notes *domain.Notes, err error) {
	if err = r.db.Where("Id = ? AND deleted = ?", id, false).First(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *notesRepository) FindNotesByTags(tags string, userId string) (notes []domain.Notes, err error) {
	if err = r.db.Where("tags = ? AND deleted = ? AND user_id = ?", tags, false, userId).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *notesRepository) FindAllNotes(userId string) (notes []domain.Notes, err error) {
	if err = r.db.Where("deleted = ? AND user_id = ?", false, userId).Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}
