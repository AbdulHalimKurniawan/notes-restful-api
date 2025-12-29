# Postman Collection & Environment

## Import ke Postman:
1. Import `Notes_API.postman_collection.json`
2. Import `Notes_API.postman_environment.json`
3. Set environment ke "Notes API Environment"

## Variables:
- `{{baseUrl}}` - Base URL API
- `{{noteId}}` - ID note yang otomatis diset setelah create

## Auto ID Management:
Setelah create note, ID otomatis tersimpan di variable `noteId` dan bisa digunakan untuk request lainnya.

---

# Postman API Testing

Base URL: `{{baseUrl}}`

## 1. Create Note
**POST** `/notes`

Headers:
```
Content-Type: application/json
```

Body (JSON):
```json
{
    "title": "My First Note",
    "content": "This is the content of my first note"
}
```

## 2. Get All Notes
**GET** `/notes`

No headers or body required.

## 3. Get Note by ID
**GET** `/notes/{{noteId}}`

No headers or body required.

## 4. Update Note
**PUT** `/notes/{{noteId}}`

Headers:
```
Content-Type: application/json
```

Body (JSON):
```json
{
    "title": "Updated Note Title",
    "content": "Updated content for the note"
}
```

## 5. Delete Note
**DELETE** `/notes/{{noteId}}`

No headers or body required.

## Expected Responses

### Success Create (201):
```json
{
    "id": 1,
    "title": "My First Note",
    "content": "This is the content of my first note",
    "created_at": "2025-12-29T23:52:43Z",
    "updated_at": "2025-12-29T23:52:43Z"
}
```

### Success Get All (200):
```json
[
    {
        "id": 1,
        "title": "My First Note",
        "content": "This is the content of my first note",
        "created_at": "2025-12-29T23:52:43Z",
        "updated_at": "2025-12-29T23:52:43Z"
    }
]
```

### Error (400/404/500):
```json
{
    "error": "Error message"
}
```