package ui_test

import (
	"encoding/json"
	"fmt"
	"github.com/gomagedon/expectate"
	_type "github.com/nicogulo/GoLearn/fullstack/backend/type"
	"github.com/nicogulo/GoLearn/fullstack/backend/ui"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Mockservice
type MockService struct {
	err   error
	todos []_type.Todo
}

func (s MockService) GetAllTodos() ([]_type.Todo, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.todos, nil
}

//Data Dummy
var dummyTodos = []_type.Todo{
	{Title: "Cooking", Description: "Cooking with mama", InCompleted: true},
	{Title: "Solve Bug", Description: "Continue Work", InCompleted: false},
}

type HTTPTest struct {
	name       string
	service    *MockService
	inputMehod string
	inputURL   string

	//expectedErr error
	expextedStatus int
	expectedTodos  []_type.Todo
}

func TestHttp(t *testing.T) {
	tests := getTest()
	tests = append(tests, getDisallowedMehodtest()...)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testHTTP(t, test)
		})
	}
}

//main test
func testHTTP(t *testing.T, test HTTPTest) {
	expect := expectate.Expect(t)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(test.inputMehod, test.inputURL, nil)
	server := ui.NewHTTP()

	server.UseService(test.service)

	server.ServerHTTP(w, r)

	var body []_type.Todo
	json.NewDecoder(w.Result().Body).Decode(&body)
	expect(w.Result().StatusCode).ToBe(test.expextedStatus)
	expect(body).ToEqual(test.expectedTodos)
}

func getTest() []HTTPTest {
	return []HTTPTest{
		{
			name:           "Random error gives 500 status and not todo",
			service:        &MockService{err: fmt.Errorf("somthing bad happened")},
			inputMehod:     "GET",
			inputURL:       "https://dev-nic.codes/todos",
			expextedStatus: 500,
			//expectedTodos: []_type.Todo,
		},
		{
			name:           "Random error gives 500 status and not todo",
			service:        &MockService{err: fmt.Errorf("somthing bad happened")},
			inputMehod:     "GET",
			inputURL:       "https://dev-nic.codes/todos/",
			expextedStatus: 500,
			//expectedTodos: []_type.Todo,
		},
		{
			name:           "wrong path gives 404 and not todos",
			service:        &MockService{todos: dummyTodos},
			inputMehod:     "GET",
			inputURL:       "https://dev-nic.codes/foo",
			expextedStatus: 404,
		},
		{
			name:           "wrong path gives 404 and not todos",
			service:        &MockService{todos: dummyTodos},
			inputMehod:     "GET",
			inputURL:       "https://dev-nic.codes/bar",
			expextedStatus: 404,
		},
		{
			name:           "return todos from service if no error",
			service:        &MockService{todos: dummyTodos},
			inputMehod:     "GET",
			inputURL:       "https://dev-nic.codes/todos/",
			expextedStatus: 200,
			expectedTodos:  dummyTodos,
		},
	}
}

func getDisallowedMehodtest() []HTTPTest {
	tests := []HTTPTest{}
	disallowedMethod := []string{
		http.MethodDelete,
		http.MethodHead,
		http.MethodPatch,
		http.MethodOptions,
		http.MethodPost,
		http.MethodPut,
	}

	for _, method := range disallowedMethod {
		tests = append(tests, HTTPTest{

			name:           fmt.Sprintf("Method %s gives 405 status and not todos", method),
			service:        &MockService{todos: dummyTodos},
			inputURL:       "https://dev-nic.codes/todos/",
			inputMehod:     method,
			expextedStatus: http.StatusMethodNotAllowed,
		})
	}
	return tests
}
