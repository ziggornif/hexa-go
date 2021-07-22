package todo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"hexa-go/infra"
	"testing"
	"time"
)

type repoMock struct{}
type repoMockErr struct{}

func (r *repoMock) Fetch() ([]Todo, error) {
	return []Todo{{
		ID:        "1234",
		Title:     "Test",
		Completed: false,
		Order:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, {
		ID:        "5678",
		Title:     "Test2",
		Completed: true,
		Order:     1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}}, nil
}

func (r *repoMock) Insert(*Todo) error {
	return nil
}

func (r *repoMock) FindByID(string) (*Todo, error) {
	return &Todo{
		ID:        "1234",
		Title:     "Test",
		Completed: false,
		Order:     0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}, nil
}

func (r *repoMock) Update(*Todo) error {
	return nil
}

func (r *repoMock) Delete(string) error {
	return nil
}

func (r *repoMockErr) Fetch() ([]Todo, error) {
	return nil, errors.New("something wrong")
}

func (r *repoMockErr) Insert(*Todo) error {
	return errors.New("something wrong")
}

func (r *repoMockErr) FindByID(id string) (*Todo, error) {
	if id == "5678" {
		return &Todo{
			ID:        "5678",
			Title:     "Test2",
			Completed: false,
			Order:     0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}, nil
	}
	return nil, errors.New("something wrong")
}

func (r *repoMockErr) Update(*Todo) error {
	return errors.New("something wrong")
}

func (r *repoMockErr) Delete(string) error {
	return errors.New("something wrong")
}

func TestListTodos(t *testing.T) {
	service := NewService(&repoMock{}, infra.GetLogger())
	todos, _ := service.ListTodos()
	assert.Equal(t, 2, len(todos))
}

func TestListTodosError(t *testing.T) {
	service := NewService(&repoMockErr{}, infra.GetLogger())
	todos, err := service.ListTodos()
	assert.Nil(t, todos)
	assert.Equal(t, err.Error.Error(), "something wrong")
}

func TestCreateTodo(t *testing.T) {
	service := NewService(&repoMock{}, infra.GetLogger())

	created, err := service.CreateTodo(Todo{
		Title:     "Test",
		Completed: false,
		Order:     0,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, created.ID)
	assert.NotEmpty(t, created.CreatedAt)
}

func TestCreateTodoError(t *testing.T) {
	service := NewService(&repoMockErr{}, infra.GetLogger())

	created, err := service.CreateTodo(Todo{
		Title:     "Test",
		Completed: false,
		Order:     0,
	})

	assert.Nil(t, created)
	assert.Equal(t, err.Error.Error(), "creation error")
}

func TestGetTodo(t *testing.T) {
	service := NewService(&repoMock{}, infra.GetLogger())

	todo, err := service.GetTodo("1234")

	assert.Nil(t, err)
	assert.Equal(t, "1234", todo.ID)
	assert.Equal(t, "Test", todo.Title)
}

func TestGetTodoError(t *testing.T) {
	service := NewService(&repoMockErr{}, infra.GetLogger())

	todo, err := service.GetTodo("1234")

	assert.Nil(t, todo)
	assert.Equal(t, err.Error.Error(), "something wrong")
}

func TestUpdateTodo(t *testing.T) {
	service := NewService(&repoMock{}, infra.GetLogger())

	todo, err := service.UpdateTodo("1234", Todo{
		Title:     "Test",
		Completed: true,
		Order:     0,
	})

	assert.Nil(t, err)
	assert.Equal(t, true, todo.Completed)
}

func TestUpdateTodoNotFountError(t *testing.T) {
	service := NewService(&repoMockErr{}, infra.GetLogger())

	todo, err := service.UpdateTodo("1234", Todo{
		Title:     "Test2",
		Completed: true,
		Order:     0,
	})

	assert.Nil(t, todo)
	assert.Equal(t, err.Error.Error(), "something wrong")
}


func TestUpdateTodoError(t *testing.T) {
	service := NewService(&repoMockErr{}, infra.GetLogger())

	todo, err := service.UpdateTodo("5678", Todo{
		Title:     "Test2",
		Completed: true,
		Order:     0,
	})

	assert.Nil(t, todo)
	assert.Equal(t, err.Error.Error(), "something wrong")
}

func TestDeleteTodo(t *testing.T) {
	service := NewService(&repoMock{}, infra.GetLogger())

	err := service.DeleteTodo("1234")

	assert.Nil(t, err)
}

func TestDeleteTodoError(t *testing.T) {
	service := NewService(&repoMockErr{}, infra.GetLogger())

	err := service.DeleteTodo("1234")

	assert.Equal(t, err.Error.Error(), "something wrong")
}
