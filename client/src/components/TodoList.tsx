import { useQuery } from "@tanstack/react-query";

import { BASE_URL } from "../constant";
import { TodoType } from "../types";
import TodoItem from "./TodoItem";
import Spinner from "./Spinner";

const TodoList = () => {
  const { data: todos, isLoading } = useQuery<TodoType[]>({
    queryKey: ["todos"],
    queryFn: async () => {
      try {
        const res = await fetch(BASE_URL + "/todos");
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.error || "Something went wrong");
        }
        console.log(data);
        return data || [];
      } catch (error) {
        console.log(error);
      }
    },
  });

  return (
    <>
      <div
        className="text-4xl uppercase font-bold text-center text-transparent bg-clip-text
        bg-gradient-to-r from-cyan-500 to-blue-500"
      >
        Todo List
      </div>

      {isLoading && (
        <div className="mx-auto">
          <Spinner />
        </div>
      )}

      <div className="flex flex-col gap-y-2">
        {todos?.map((todo) => (
          <TodoItem key={todo._id} todo={todo} />
        ))}
      </div>
    </>
  );
};

export default TodoList;
