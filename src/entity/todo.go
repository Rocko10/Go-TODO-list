package entity

type Todo struct {
	id   int
	Name string
	Done bool
}

func NewTodo(name string) Todo {
	return Todo{id: -1, Name: name, Done: false}
}

func (t Todo) GetId() int {
	return t.id
}

func (t *Todo) SetId(id int) {
	t.id = id
}
