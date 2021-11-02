package ui

import _type "github.com/nicogulo/todo-app-go-react/backend/type"

type Service interface {
	GetAllTodos() ([]_type.Todo, error)
}
