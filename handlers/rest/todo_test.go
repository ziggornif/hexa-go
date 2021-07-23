package rest

import (
	"encoding/json"
	"hexa-go/infra/config"
	"hexa-go/infra/logger"
	"hexa-go/infra/storage"
	"hexa-go/packages/todo"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
)

var db *pg.DB

func cleanInserted(id string) {
	db.Model(&todo.Todo{}).Where("id = ?", id).Delete()
}

func init() {
	conf, _ := config.LoadConfig("../..", logger.GetLogger())
	db, _ = storage.DBConnect(conf.GetConfig(), logger.GetLogger())

	router := gin.Default()
	apiRouter := router.Group("/api")

	repo, _ := todo.NewRepository(db, logger.GetLogger())
	service := todo.NewService(repo, logger.GetLogger())
	NewTodoController(apiRouter, service, logger.GetLogger())

	go router.Run()
}

func TestPostTodo(t *testing.T) {
	r := strings.NewReader(`
	{
		"title": "task",
		"order": 0,
		"completed": false
	}
	`)
	resp, _ := http.Post("http://localhost:8080/api/todos/",
		"application/json", r)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var entity todo.Todo
	json.Unmarshal(data, &entity)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "task", entity.Title)

	cleanInserted(entity.ID)
}

func TestPostTodoError(t *testing.T) {
	r := strings.NewReader(`
	{
		"title": "task",
		"order": 0,
		"completed": false
	}
	`)
	resp, _ := http.Post("http://localhost:8080/api/todos/",
		"application/json", r)
	data, _ := ioutil.ReadAll(resp.Body)
	var entity todo.Todo
	json.Unmarshal(data, &entity)
	assert.Equal(t, 200, resp.StatusCode)
	resp.Body.Close()

	r2 := strings.NewReader(`
	{
		"title": "task",
		"order": 0,
		"completed": false
	}
	`)
	resp2, _ := http.Post("http://localhost:8080/api/todos",
		"application/json", r2)
	var error map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&error)

	assert.Equal(t, "Error - creation error", error["error"].(string))
	assert.Equal(t, 500, resp2.StatusCode)
	resp2.Body.Close()

	cleanInserted(entity.ID)
}

func TestPostTodoBadRequestError(t *testing.T) {
	resp, _ := http.Post("http://localhost:8080/api/todos/",
		"application/json", nil)
	var error map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&error)

	assert.Equal(t, "Invalid entity", error["error"].(string))
	assert.Equal(t, 400, resp.StatusCode)
	resp.Body.Close()
}
