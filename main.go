package main

import (
	"log"
	"notes-api/config"
	"notes-api/internal/handler"
	"notes-api/internal/repository"
	"notes-api/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Dependency injection
	noteRepo := repository.NewMysqlNoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NewNoteHandler(noteService)

	// Setup routes
	r := gin.Default()
	
	api := r.Group("/api/v1")
	{
		api.POST("/notes", noteHandler.CreateNote)
		api.GET("/notes", noteHandler.GetAllNotes)
		api.GET("/notes/:id", noteHandler.GetNoteByID)
		api.PUT("/notes/:id", noteHandler.UpdateNote)
		api.DELETE("/notes/:id", noteHandler.DeleteNote)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}