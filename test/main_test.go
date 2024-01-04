package test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"simple-to-do/app/api"
	"simple-to-do/app/domains"
	"simple-to-do/app/repositories"
	"simple-to-do/app/usecases"
	"simple-to-do/config"
	"strconv"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupDatasource() *sql.DB {
	return config.InitMysqlMasterData()
}
func setupApp() http.Handler {
	d := setupDatasource()
	v := validator.New()

	repo := repositories.NewDatamaster()
	usecase := usecases.NewTodoUsecase(repo, d, v)
	router := api.NewApi(usecase)
	handler := config.Router(router)
	return handler
}

func setupTruncate(db *sql.DB) {
	db.Exec("TRUNCATE task")
}

func TestGetAllItemsNull(t *testing.T) {
	setupTruncate(setupDatasource())
	router := setupApp()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/tasks", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusOK, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), resBody["status"])
	assert.Nil(t, resBody["error"])
	assert.Nil(t, resBody["data"])
}

func TestGetAllItems(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	tx, _ := d.Begin()
	repo := repositories.NewDatamaster()
	repo.Save(context.Background(), tx, domains.Task{
		Title:       "Self development",
		Description: "Reading 10 pages of a programming book",
		DueDate:     time.Now().Add(2 * 24 * time.Hour),
	})
	tx.Commit()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/tasks", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusOK, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), resBody["status"])
	assert.Nil(t, resBody["error"])
	assert.NotNil(t, resBody["data"])
}

func TestGetItemSuccess(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	tx, _ := d.Begin()
	repo := repositories.NewDatamaster()
	item, _ := repo.Save(context.Background(), tx, domains.Task{
		Title:       "Self development",
		Description: "Reading 10 pages of a programming book",
		DueDate:     time.Now().Add(2 * 24 * time.Hour),
	})
	tx.Commit()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/tasks/"+strconv.Itoa(item.Id), nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusOK, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), resBody["status"])
	assert.Nil(t, resBody["error"])
	assert.NotNil(t, resBody["data"])
	assert.Equal(t, item.Id, int(resBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, item.Title, resBody["data"].(map[string]interface{})["title"])
}

func TestGetItemFailed(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	tx, _ := d.Begin()
	repo := repositories.NewDatamaster()
	repo.Save(context.Background(), tx, domains.Task{
		Title:       "Self development",
		Description: "Reading 10 pages of a programming book",
		DueDate:     time.Now().Add(2 * 24 * time.Hour),
	})
	tx.Commit()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/tasks/999", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusNotFound, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), resBody["status"])
	assert.Equal(t, http.StatusText(http.StatusNotFound), resBody["error"])
	assert.Nil(t, resBody["data"])
}

func TestCreateItemSuccess(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	task := domains.Task{
		Title:       "Self development",
		Description: "Reading 10 pages of a programming book",
		DueDate:     time.Now().Add(2 * 24 * time.Hour),
	}
	jsonData, _ := json.Marshal(task)
	reqBody := bytes.NewReader(jsonData)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/tasks", reqBody)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusCreated, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusCreated), resBody["status"])
	assert.Nil(t, resBody["error"])
	assert.Equal(t, task.Title, resBody["data"].(map[string]interface{})["title"])
}

func TestCreateItemFailed(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	task := domains.Task{
		Description: "Reading 10 pages of a programming book",
		DueDate:     time.Now().Add(2 * 24 * time.Hour),
	}
	jsonData, _ := json.Marshal(task)
	reqBody := bytes.NewReader(jsonData)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/tasks", reqBody)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusBadRequest, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusBadRequest), resBody["status"])
	assert.NotNil(t, resBody["error"])
	assert.Nil(t, resBody["data"])
}

func TestUpdateItemSuccess(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	task := domains.Task{
		Title:       "Self development",
		Description: "Reading 10 pages of a programming book",
		DueDate:     time.Now().Add(2 * 24 * time.Hour),
	}

	tx, _ := d.Begin()
	repo := repositories.NewDatamaster()
	item, _ := repo.Save(context.Background(), tx, task)
	tx.Commit()

	task.Title = "Programming"

	jsonData, _ := json.Marshal(task)
	reqBody := bytes.NewReader(jsonData)
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/tasks/"+strconv.Itoa(item.Id), reqBody)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusAccepted, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusAccepted, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusAccepted), resBody["status"])
	assert.Nil(t, resBody["error"])
	assert.Equal(t, item.Id, int(resBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, task.Title, resBody["data"].(map[string]interface{})["title"])
}

func TestUpdateItemFailed(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	task := domains.Task{
		Title:       "Self development",
		Description: "Reading 10 pages of a programming book",
		DueDate:     time.Now().Add(2 * 24 * time.Hour),
	}

	tx, _ := d.Begin()
	repo := repositories.NewDatamaster()
	item, _ := repo.Save(context.Background(), tx, task)
	tx.Commit()

	task.Title = ""

	jsonData, _ := json.Marshal(task)
	reqBody := bytes.NewReader(jsonData)
	req := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/tasks/"+strconv.Itoa(item.Id), reqBody)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusBadRequest, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusBadRequest), resBody["status"])
	assert.NotNil(t, resBody["error"])
	assert.Nil(t, resBody["data"])
}
func TestDeleteItemSuccess(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	tx, _ := d.Begin()
	repo := repositories.NewDatamaster()
	item, _ := repo.Save(context.Background(), tx, domains.Task{
		Title:       "Self development",
		Description: "Reading 10 pages of a programming book",
		DueDate:     time.Now().Add(2 * 24 * time.Hour),
	})
	tx.Commit()

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/tasks/"+strconv.Itoa(item.Id), nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusOK, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), resBody["status"])
	assert.Nil(t, resBody["error"])
	assert.Nil(t, resBody["data"])
}

func TestDeleteItemFailed(t *testing.T) {
	d := setupDatasource()
	setupTruncate(d)
	router := setupApp()

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/tasks/999", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	res := rec.Result()
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	body, _ := io.ReadAll(res.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, http.StatusNotFound, int(resBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), resBody["status"])
	assert.Equal(t, http.StatusText(http.StatusNotFound), resBody["error"])
	assert.Nil(t, resBody["data"])
}
