package gateway

import (
	"errors"

	"github.com/Rocko10/Go-TODO-list/src/entity"
)

type TodoRepo struct {
	todos []entity.Todo
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{}
}

func (tr TodoRepo) GetAll() []entity.Todo {
	return tr.todos
}

func (tr *TodoRepo) Add(todo entity.Todo) {
	id := len(tr.todos) + 1
	todo.SetId(id)

	tr.todos = append(tr.todos, todo)
}

func (tr TodoRepo) Get(id int) (*entity.Todo, error) {
	t := entity.Todo{}

	for i, _ := range tr.todos {
		if tr.todos[i].GetId() == id {
			return &(tr.todos[i]), nil
		}
	}

	return &t, errors.New("Not founded Todo with that id")
}

func (tr *TodoRepo) Toggle(id int) (entity.Todo, error) {
	t := entity.Todo{}

	for i, _ := range tr.todos {
		if tr.todos[i].GetId() == id {
			tr.todos[i].Done = !tr.todos[i].Done

			return tr.todos[i], nil
		}
	}

	return t, errors.New("Not founed Todo with that id")
}

func (tr *TodoRepo) Delete(id int) error {
	for i := range tr.todos {
		if tr.todos[i].GetId() == id {
			tr.todos = append(tr.todos[:i], tr.todos[i+1:]...)

			return nil
		}
	}

	return errors.New("No Todo founded with that id")
}
