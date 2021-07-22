package todo

import (
	"errors"
	error "hexa-go/infra/error"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// Service - service interface
type Service interface {
	ListTodos() ([]Todo, *error.HexagoError)
	CreateTodo(todo Todo) (*Todo, *error.HexagoError)
	GetTodo(id string) (*Todo, *error.HexagoError)
	UpdateTodo(id string, todo Todo) (*Todo, *error.HexagoError)
	DeleteTodo(id string) *error.HexagoError
}

type service struct {
	repo   Repository
	logger *logrus.Logger
}

// NewService - create service instance
func NewService(repo Repository, logger *logrus.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

// ListTodos - list todos
func (s *service) ListTodos() ([]Todo, *error.HexagoError) {
	todos, err := s.repo.Fetch()
	if err != nil {
		s.logger.Errorf("[Service] ListTodos - Error while getting todos")
		s.logger.Debugf("... context : %s", err)
		return nil, &error.HexagoError{Error: err}
	}

	return todos, nil
}

// CreateTodo - create todo
func (s *service) CreateTodo(todo Todo) (*Todo, *error.HexagoError) {
	id, _ := uuid.NewV4()
	todo.ID = id.String()
	todo.CreatedAt = time.Now()

	err := s.repo.Insert(&todo)
	if err != nil {
		s.logger.Errorf("[Service] CreateTodo - Error while creating todo")
		s.logger.Debugf("... context : %s", err)
		return nil, &error.HexagoError{Error: errors.New("creation error")}
	}

	return &todo, nil
}

// GetTodo - get todo by id
func (s *service) GetTodo(id string) (*Todo, *error.HexagoError) {
	todo, err := s.repo.FindByID(id)
	if err != nil {
		s.logger.Errorf("[Service] GetTodo - Error while getting todo %s", id)
		s.logger.Debugf("... context : %s", err)
		return nil, &error.HexagoError{Kind: "NotFound", Error: errors.New("todo not found")}
	}

	return todo, nil
}

// UpdateTodo - update todo
func (s *service) UpdateTodo(id string, update Todo) (*Todo, *error.HexagoError) {
	todo, getErr := s.GetTodo(id)
	if getErr != nil {
		return nil, getErr
	}

	todo.Title = update.Title
	todo.Order = update.Order
	todo.Completed = update.Completed
	todo.UpdatedAt = time.Now()

	err := s.repo.Update(todo)
	if err != nil {
		s.logger.Errorf("[Service] UpdateTodo - Error while updating todo")
		s.logger.Debugf("... context : %s", err)
		return nil, &error.HexagoError{Error: err}
	}

	return todo, nil
}

// DeleteTodo - delete todo
func (s *service) DeleteTodo(id string) *error.HexagoError {
	err := s.repo.Delete(id)
	if err != nil {
		s.logger.Errorf("[Service] DeleteTodo - Error while removing todo %s", id)
		s.logger.Debugf("... context : %s", err)
		return &error.HexagoError{Error: err}
	}

	s.logger.Info("[Service] DeleteTodo - todo item removed")
	return nil
}
