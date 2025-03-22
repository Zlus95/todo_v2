import React, { useState, useEffect } from "react";
import { Header } from "../Header/Header";
import AddTodo from "./AddTodo";
import { useQuery } from "@tanstack/react-query";
import api from "../api";
import Todo from "./Todo";

function useGetTodos() {
  return useQuery({
    queryKey: ["todoList"],
    queryFn: async () => {
      const response = await api.get("/tasks");
      return response;
    },
  });
}

const TodoPage = () => {
  const [todo, setTodo] = useState([]);
  const { isLoading, data, error, isError, isSuccess } = useGetTodos();

  useEffect(() => {
    if (isSuccess) {
      setTodo(data);
    }
  }, [data, isSuccess]);

  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-full text-primary text-2xl">
        Loading...
      </div>
    );
  }

  if (isError) {
    return (
      <div className="flex justify-center items-center h-full text-red-500 text-2xl">
        Error: {error.message}
      </div>
    );
  }

  return (
    <>
      <div className="h-full w-full">
        <Header />
        <div className="flex justify-center mt-6">
          <AddTodo />
        </div>
        <div className="flex justify-center flex-col gap-4 items-center mt-6 h-96 overflow-y-scroll">
          {(todo.data || []).map((item) => (
            <div
              key={item.id}
              className="w-64 border-2 border-orange-300 flex justify-between p-2 rounded"
            >
              <Todo todo={todo} {...item} />
            </div>
          ))}
        </div>
      </div>
    </>
  );
};

export default TodoPage;
