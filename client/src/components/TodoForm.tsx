import { useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { IoMdAdd } from "react-icons/io";

import { BASE_URL } from "../constant";
import Spinner from "./Spinner";

const TodoForm = () => {
  const [newTodo, setNewTodo] = useState("");

  const queryClient = useQueryClient();

  const { mutate: createTodo, isPending: isCreating } = useMutation({
    mutationKey: ["createTodo"],
    mutationFn: async (e: React.FormEvent) => {
      e.preventDefault();
      try {
        const res = await fetch(BASE_URL + `/todos`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ body: newTodo }),
        });
        const data = await res.json();

        if (!res.ok) {
          throw new Error(data.error || "Something went wrong");
        }

        setNewTodo("");
        return data;
      } catch (error: unknown) {
        throw new Error((error as Error).message);
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
    onError: (error: Error) => {
      alert(error.message);
    },
  });

  return (
    <form className="w-full" onSubmit={createTodo}>
      <div className="flex items-center gap-x-3 w-full">
        <input
          type="text"
          className="w-full rounded-lg bg-white py-2 px-3 text-black text-lg"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          ref={(input) => input && input.focus()}
        />

        <button
          className="bg-gray-700 rounded-lg py-2 px-3 text-white"
          type="submit"
        >
          {isCreating ? <Spinner /> : <IoMdAdd size={30} />}
        </button>
      </div>
    </form>
  );
};

export default TodoForm;
