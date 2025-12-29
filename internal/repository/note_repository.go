package repository

import "notes-api/internal/model"

type NoteRepository interface {
	Create(note *model.Note) error
	GetAll() ([]model.Note, error)
	GetByID(id int) (*model.Note, error)
	Update(id int, note *model.Note) error
	Delete(id int) error
}