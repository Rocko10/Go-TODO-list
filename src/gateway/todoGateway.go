package gateway

import "github.com/Rocko10/Go-TODO-list/src/entity"

type TodoGateway interface {
	GetAll() []entity.Todo
	Add(todo entity.Todo)
	Get(id int) (*entity.Todo, error)
	Toggle(id int) (entity.Todo, error)
	Delete(id int) error
}
