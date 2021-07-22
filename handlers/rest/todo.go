package rest

import (
	resterror "hexa-go/handlers/rest/error"
	"hexa-go/packages/todo"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

// TodoController - controller interface
type TodoController interface {
	postTodo(c *gin.Context)
	getTodos(c *gin.Context)
	getTodo(c *gin.Context)
	updateTodo(c *gin.Context)
	deleteTodo(c *gin.Context)
}

type tcontroller struct {
	service todo.Service
	logger  *logrus.Logger
}

// NewTodoController - create controller instance
func NewTodoController(route *gin.RouterGroup, service todo.Service, logger *logrus.Logger) {
	controller := tcontroller{
		service: service,
		logger:  logger,
	}

	todoRouter := route.Group("/todos")
	{
		todoRouter.POST("/", controller.postTodo)
		todoRouter.GET("/", controller.getTodos)
		todoRouter.GET("/:id", controller.getTodo)
		todoRouter.PUT("/:id", controller.updateTodo)
		todoRouter.DELETE("/:id", controller.deleteTodo)
	}
}

// postTodo - create todo
func (ctrl *tcontroller) postTodo(c *gin.Context) {
	query := todo.Todo{}
	if err := c.ShouldBindBodyWith(&query, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := ctrl.service.CreateTodo(query)
	if err != nil {
		resterror.SendHTTPError(c, err)
		return
	}

	c.JSON(http.StatusOK, entity)
}

// getTodos - get todo list
func (ctrl *tcontroller) getTodos(c *gin.Context) {
	todos, err := ctrl.service.ListTodos()
	if err != nil {
		resterror.SendHTTPError(c, err)
		return
	}

	c.JSON(http.StatusOK, todos)
}

// getTodo - get todo item
func (ctrl *tcontroller) getTodo(c *gin.Context) {
	todoID := c.Param("id")
	entity, err := ctrl.service.GetTodo(todoID)
	if err != nil {
		resterror.SendHTTPError(c, err)
		return
	}

	c.JSON(http.StatusOK, entity)
}

// updateTodo - update todo item
func (ctrl *tcontroller) updateTodo(c *gin.Context) {
	todoID := c.Param("id")
	query := todo.Todo{}
	if err := c.ShouldBindBodyWith(&query, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTodo, err := ctrl.service.UpdateTodo(todoID, query)

	if err != nil {
		resterror.SendHTTPError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

// deleteTodo - delete todo item
func (ctrl *tcontroller) deleteTodo(c *gin.Context) {
	todoID := c.Param("id")

	err := ctrl.service.DeleteTodo(todoID)
	if err != nil {
		resterror.SendHTTPError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
