package service

import (
	"github.com/Rocko10/Go-TODO-list/src/entity"
	"github.com/Rocko10/Go-TODO-list/src/gateway"
)

type TodoService struct {
	todoGateway gateway.TodoGateway
}

func NewTodoService(todoGateway gateway.TodoGateway) TodoService {
	return TodoService{todoGateway: todoGateway}
}

func (ts TodoService) GetAll() []entity.Todo {
	return ts.todoGateway.GetAll()
}

func (ts TodoService) Add(todo entity.Todo) {
	ts.todoGateway.Add(todo)
}

func (ts TodoService) Get(id int) (*entity.Todo, error) {
	return ts.todoGateway.Get(id)
}

func (ts TodoService) Toggle(id int) (entity.Todo, error) {
	return ts.todoGateway.Toggle(id)
}

func (ts TodoService) Delete(id int) error {
	return ts.todoGateway.Delete(id)
}
