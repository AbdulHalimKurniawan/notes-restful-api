package service

import (
	"notes-api/internal/model"
	"notes-api/internal/repository"
)

type NoteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(noteRepo repository.NoteRepository) *NoteService {
	return &NoteService{noteRepo: noteRepo}
}

func (s *NoteService) CreateNote(req *model.CreateNoteRequest) (*model.Note, error) {
	note := &model.Note{
		Title:   req.Title,
		Content: req.Content,
	}
	err := s.noteRepo.Create(note)
	return note, err
}

func (s *NoteService) GetAllNotes() ([]model.Note, error) {
	return s.noteRepo.GetAll()
}

func (s *NoteService) GetNoteByID(id int) (*model.Note, error) {
	return s.noteRepo.GetByID(id)
}

func (s *NoteService) UpdateNote(id int, req *model.UpdateNoteRequest) error {
	note := &model.Note{
		Title:   req.Title,
		Content: req.Content,
	}
	return s.noteRepo.Update(id, note)
}

func (s *NoteService) DeleteNote(id int) error {
	return s.noteRepo.Delete(id)
}