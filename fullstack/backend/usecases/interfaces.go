package usecases

import _type "github.com/nicogulo/GoLearn/fullstack/backend/type"

type TodosRepository interface {
	GetAllTodos() ([]_type.Todo, error)
}
