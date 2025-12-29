package repository

import (
	"database/sql"
	"notes-api/internal/model"
	"time"
)

type postgresNoteRepository struct {
	db *sql.DB
}

func NewPostgresNoteRepository(db *sql.DB) NoteRepository {
	return &postgresNoteRepository{db: db}
}

func (r *postgresNoteRepository) Create(note *model.Note) error {
	query := "INSERT INTO notes (title, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	now := time.Now()
	var id int
	err := r.db.QueryRow(query, note.Title, note.Content, now, now).Scan(&id)
	if err != nil {
		return err
	}
	note.ID = id
	note.CreatedAt = now
	note.UpdatedAt = now
	return nil
}

func (r *postgresNoteRepository) GetAll() ([]model.Note, error) {
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

func (r *postgresNoteRepository) GetByID(id int) (*model.Note, error) {
	query := "SELECT id, title, content, created_at, updated_at FROM notes WHERE id = $1"
	var note model.Note
	err := r.db.QueryRow(query, id).Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *postgresNoteRepository) Update(id int, note *model.Note) error {
	query := "UPDATE notes SET title = $1, content = $2, updated_at = $3 WHERE id = $4"
	now := time.Now()
	_, err := r.db.Exec(query, note.Title, note.Content, now, id)
	return err
}

func (r *postgresNoteRepository) Delete(id int) error {
	query := "DELETE FROM notes WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}