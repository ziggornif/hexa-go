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
	resp, _ := http.Post("http://localhost:8080/api/todos",
		"application/json", r)
	data, _ := ioutil.ReadAll(resp.Body)
	var entity todo.Todo
	json.Unmarshal(data, &entity)

	assert.Equal(t, 200, resp.StatusCode)

	cleanInserted(entity.ID)
}
