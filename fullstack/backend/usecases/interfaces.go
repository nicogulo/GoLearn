package usecases

import _type "github.com/nicogulo/todo-app-go-react/backend/type"

type TodosRepository interface {
	 GetAllTodos() ([] _type.Todo, error)
 }