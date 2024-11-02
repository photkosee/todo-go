import { useMutation, useQueryClient } from "@tanstack/react-query";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";

import { BASE_URL } from "../constant";
import { TodoType } from "../types";
import Spinner from "./Spinner";

const TodoItem = ({ todo }: { todo: TodoType }) => {
  const queryClient = useQueryClient();

  const { mutate: updateTodo, isPending: isUpdating } = useMutation({
    mutationKey: ["updateTodo"],
    mutationFn: async () => {
      if (todo.completed) return alert("Todo is already completed");
      console.log(todo._id);
      try {
        const res = await fetch(BASE_URL + `/todos/${todo._id}/complete`, {
          method: "PATCH",
        });
        const data = await res.json();
        if (!res.ok) {
          throw new Error(data.error || "Something went wrong");
        }
        return data;
      } catch (error) {
        console.log(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
    mutationKey: ["deleteTodo"],
    mutationFn: async () => {
      try {
        const res = await fetch(BASE_URL + `/todos/${todo._id}`, {
          method: "DELETE",
        });
        const data = await res.json();
        if (!res.ok) {
          throw new Error(data.error || "Something went wrong");
        }
        return data;
      } catch (error) {
        console.log(error);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  return (
    <div className="flex gap-x-2 items-center w-full justify-between">
      <div
        className="flex justify-between items-center py-2 px-3 rounded-lg
        border border-gray-500 w-full"
      >
        <div className="flex justify-between">
          <div
            className={`${
              todo.completed ? "line-through text-green-500" : "text-yellow-500"
            }`}
          >
            {todo.body}
          </div>
        </div>

        {todo.completed && (
          <div className="text-green-500 bg-green-900 px-1 rounded-sm font-bold">
            Done
          </div>
        )}
        {!todo.completed && (
          <div className="text-yellow-500 bg-yellow-900 px-1 rounded-sm font-bold">
            In Progress
          </div>
        )}
      </div>

      <div className="flex gap-x-1 items-center">
        {!todo.completed && (
          <button onClick={() => updateTodo()}>
            {!isUpdating && (
              <FaCheckCircle size={20} className="text-green-500" />
            )}
            {isUpdating && <Spinner />}
          </button>
        )}

        <button onClick={() => deleteTodo()}>
          {!isDeleting && <MdDelete size={25} className="text-red-500" />}
          {isDeleting && <Spinner />}
        </button>
      </div>
    </div>
  );
};

export default TodoItem;
