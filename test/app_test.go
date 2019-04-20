package test

import (
	"testing"

	"github.com/Rocko10/Go-TODO-list/src/entity"
	"github.com/Rocko10/Go-TODO-list/src/gateway"
	"github.com/Rocko10/Go-TODO-list/src/service"
)

func TestTodoStruct(t *testing.T) {
	todo := entity.NewTodo("Wash dishes")

	t.Log("Todo created:")
	t.Log(todo)
}

func TestTodoService(t *testing.T) {
	var todoRepo gateway.TodoGateway = gateway.NewTodoRepo()
	var todoService service.TodoService = service.NewTodoService(todoRepo)
	var todo entity.Todo = entity.NewTodo("Wash dishes")

	todoService.Add(todo)
	selectedTodo, err := todoService.Get(100)
	toggledTodo, err2 := todoService.Toggle(100)
	err3 := todoService.Delete(100)

	if err == nil || err2 == nil || err3 == nil {
		t.Fatal("Should throw, since no todo with that id exist")
	} else {
		t.Logf("The expected error (1) is: %s", err)
		t.Logf("The expected error (2) is: %s", err2)
		t.Logf("The expected error (3) is: %s", err3)
	}

	selectedTodo, err = todoService.Get(1)
	toggledTodo, err2 = todoService.Toggle(1)

	if err != nil || err2 != nil {
		t.Fatal("Should not throw, since todo with id 1 exists")
	}

	var todos []entity.Todo = todoService.GetAll()

	if todos[0].Done == false {
		t.Log("GetAll todos value is:")
		t.Log(todos)
		t.Fatal("Should be true, since past value was false")
	}

	t.Log("GetAll todos value is:")
	t.Log(todos)
	t.Log("Get 1 todo:")
	t.Log(*selectedTodo)
	t.Log("Toggle todo is:")
	t.Log(toggledTodo)

	t.Log("Deleting todo with id 1...")
	err3 = todoService.Delete(1)

	if err3 != nil {
		t.Fatal("Should be deleted the todo, since exists")
	}

	todos = todoService.GetAll()
	t.Log("After delete 1, the todos are:")
	t.Log(todos)
}
