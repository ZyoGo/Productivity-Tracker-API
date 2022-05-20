package todos

import (
	domain "github.com/w33h/Productivity-Tracker-API/business/todos"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.RepositoryTodos {
	return &todoRepository{db}
}

func (r *todoRepository) InsertTodo(todo domain.Todo) (id string, err error) {
	if err = r.db.Create(&todo).Error; err != nil {
		return id, err
	}

	id = todo.Id

	return id, nil
}

func (r *todoRepository) UpdateTodo(todo domain.Todo) (err error) {
	if err = r.db.Save(&todo).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepository) DeleteTodo(id string) (err error) {
	if err = r.db.Table("todos").Where("Id = ?", id).Update("deleted", true).Error; err != nil {
		return err
	}

	return nil
}

func (r *todoRepository) FindByStatus(status string) (todo []domain.Todo, err error) {
	if err = r.db.Where("status = ? AND deleted = ?", status, false).Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, err
}

func (r *todoRepository) FindById(id string) (todo *domain.Todo, err error) {
	if err = r.db.Where("Id = ? AND deleted = ?", id, false).First(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *todoRepository) FindAllTodo(userId string) (todo []domain.Todo, err error) {
	if err = r.db.Where("deleted = ? AND user_id = ?", false, userId).Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
