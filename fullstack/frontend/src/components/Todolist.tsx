import React from "react";
import { TodoItem } from "./TodoItem";
import { Todo } from "./types/Todo";
type Props = {
  todos: Todo[];
};
const Todolist: React.FC<Props> = ({ todos }) => {
  return (
    <div className="todo-list">
      <ul>
        {todos.map((todo, i) => (
          <li key={i}>
            <TodoItem
              title={todo.title}
              desctopstion={todo.desctopstion}
              inCompleted={todo.inCompleted}
            />
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Todolist;
