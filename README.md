# Notes RESTful API

RESTful API untuk aplikasi Notes menggunakan Golang, Gin framework, dan MySQL dengan Clean Architecture.

## Struktur Proyek

```
notes-api/
├── config/
│   └── database.go          # Konfigurasi database
├── database/
│   └── migration.sql        # SQL migration
├── internal/
│   ├── handler/
│   │   └── note_handler.go  # HTTP handlers
│   ├── model/
│   │   └── note.go          # Data models
│   ├── repository/
│   │   ├── note_repository.go       # Repository interface
│   │   └── mysql_note_repository.go # MySQL implementation
│   └── service/
│       └── note_service.go  # Business logic
├── .env                     # Environment variables
├── go.mod                   # Go modules
└── main.go                  # Entry point
```

## Endpoints

- `POST /api/v1/notes` - Create note
- `GET /api/v1/notes` - Get all notes
- `GET /api/v1/notes/{id}` - Get note by ID
- `PUT /api/v1/notes/{id}` - Update note
- `DELETE /api/v1/notes/{id}` - Delete note

## Alur Request

Request → Handler → Service → Repository → Database

## Setup

1. Install dependencies: `go mod tidy`
2. Setup MySQL database
3. Run migration: `mysql < database/migration.sql`
4. Update .env file
5. Run: `go run main.go`