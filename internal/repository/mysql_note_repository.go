package repository

import (
	"database/sql"
	"notes-api/internal/model"
	"time"
)

type mysqlNoteRepository struct {
	db *sql.DB
}

func NewMysqlNoteRepository(db *sql.DB) NoteRepository {
	return &mysqlNoteRepository{db: db}
}

func (r *mysqlNoteRepository) Create(note *model.Note) error {
	query := "INSERT INTO notes (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)"
	now := time.Now()
	result, err := r.db.Exec(query, note.Title, note.Content, now, now)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	note.ID = int(id)
	note.CreatedAt = now
	note.UpdatedAt = now
	return nil
}

func (r *mysqlNoteRepository) GetAll() ([]model.Note, error) {
	query := "SELECT id, title, content, created_at, updated_at FROM notes"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func (r *mysqlNoteRepository) GetByID(id int) (*model.Note, error) {
	query := "SELECT id, title, content, created_at, updated_at FROM notes WHERE id = ?"
	var note model.Note
	err := r.db.QueryRow(query, id).Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *mysqlNoteRepository) Update(id int, note *model.Note) error {
	query := "UPDATE notes SET title = ?, content = ?, updated_at = ? WHERE id = ?"
	now := time.Now()
	_, err := r.db.Exec(query, note.Title, note.Content, now, id)
	return err
}

func (r *mysqlNoteRepository) Delete(id int) error {
	query := "DELETE FROM notes WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}