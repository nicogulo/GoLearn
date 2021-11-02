package usecases

import _type "github.com/nicogulo/GoLearn/fullstack/backend/type"

func GetTodos(repo TodosRepository) ([]_type.Todo, error) {
	todos, err := repo.GetAllTodos()
	if err != nil {
		return nil, ErrInternal
	}
	return todos, nil
}
