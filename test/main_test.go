package test

import (
	"database/sql"
	"fairusatoir/simple-to-do/todo"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRouter() http.Handler {
	return todo.SetRouter()
}

func truncate(db *sql.DB) {
	db.Exec("TRUNCATE task")
}

// func TestSaveTest(t *testing.T) {
// 	db := todo.SetPool()
// 	truncate(db)

// 	router := setupRouter()

// 	reqBody := strings.NewReader(`{"title":"Self development","description":"Reading 10 pages of a basic book","due_date":"2023-03-16T15:04:05.999999-07:00"}`)
// 	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/tasks", reqBody)
// 	fmt.Print(req)
// 	req.Header.Add("Content-Type", "application/json")

// 	rec := httptest.NewRecorder()

// 	router.ServeHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)

// }

func TestGetAllItems(t *testing.T) {
	router := setupRouter()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/tasks", nil)
	req.Header.Add("Content-Type", "application/json")
	fmt.Print(req)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
