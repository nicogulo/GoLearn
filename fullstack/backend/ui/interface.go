package ui

import _type "github.com/nicogulo/GoLearn/fullstack/backend/type"

type Service interface {
	GetAllTodos() ([]_type.Todo, error)
}
