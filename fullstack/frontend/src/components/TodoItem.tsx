import React from "react";
import { Todo } from "./types/Todo";

export const TodoItem: React.FC<Todo> = ({
  title,
  desctopstion,
  inCompleted,
}) => {
  return (
    <article className="todo-item">
      <h3>{title}</h3>
      <p>{desctopstion}</p>
      <input type="checkbox" defaultChecked={inCompleted} />
    </article>
  );
};
