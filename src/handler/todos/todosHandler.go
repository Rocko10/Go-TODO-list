package todos

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Rocko10/Go-TODO-list/src/entity"
	"github.com/Rocko10/Go-TODO-list/src/gateway"
	"github.com/Rocko10/Go-TODO-list/src/service"
)

type listData struct {
	Todos []entity.Todo
}

var todoRepo gateway.TodoGateway = gateway.NewTodoRepo()
var todoService service.TodoService = service.NewTodoService(todoRepo)

func init() {
	log.Println("[init] on todos handler called...")
}

func List(w http.ResponseWriter, r *http.Request) {
	listView := template.Must(template.ParseFiles("views/todos/list.html"))
	todos := todoService.GetAll()
	data := listData{Todos: todos}

	listView.Execute(w, data)
}

func New(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	if len(name) > 0 {
		todo := entity.NewTodo(name)
		todoService.Add(todo)
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Toggle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("todo"))

	todoService.Toggle(id)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("todo"))

	todoService.Delete(id)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
