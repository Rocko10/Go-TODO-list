package todos

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Rocko10/todoApp/src/entity"
	"github.com/Rocko10/todoApp/src/gateway"
	"github.com/Rocko10/todoApp/src/service"
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
