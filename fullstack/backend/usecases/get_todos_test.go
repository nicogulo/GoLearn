package usecases_test

import (
	"fmt"
	"github.com/gomagedon/expectate"
	_type "github.com/nicogulo/todo-app-go-react/backend/type"
	"github.com/nicogulo/todo-app-go-react/backend/usecases"
	"testing"
)

var dummyTodos = []_type.Todo{
	{Title: "Cooking", Description: "Cooking with mama", InCompleted: true},
	{Title: "Solve Bug", Description: "Continue Work", InCompleted: false},
}

type BadTodosrepo struct {

}

func (BadTodosrepo) GetAllTodos() ([]_type.Todo, error) {
	return nil, fmt.Errorf("Something went wrong")
}

type MockTodoRepo struct {

}
func (MockTodoRepo) GetAllTodos() ([]_type.Todo, error) {
	//return nil, fmt.Errorf("Something went wrong")
	return dummyTodos, nil
}

func TestGetTodos( t *testing.T)  {
	t.Run("Return ErrInternal when TodosRepositoty return error ", func(t *testing.T) {
		expect := expectate.Expect(t)

		repo := new(BadTodosrepo)
		todos, err := usecases.GetTodos(repo)

		expect(err).ToBe(usecases.ErrInternal)
		//expect(todos).ToBe(nill)
		//if err != usecases.ErrInternal{
		//	t.Fatalf("Expected ErrInternal; Got: %v", err)
		//}
		if todos != nil {
			//t.Fatalf("Expected todos be nill; Got: %v", todos)
			t.Fatalf("Expected todos to be nil")
		}
	})

	t.Run("Return todos from TodosReposity", func(t *testing.T) {
	expect := expectate.Expect(t)

	repo := new(MockTodoRepo)
	todos, err := usecases.GetTodos(repo)

	expect(err).ToBe(nil)
	expect(todos).ToEqual(dummyTodos)
	})

}