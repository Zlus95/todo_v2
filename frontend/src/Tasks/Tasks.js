import React, { useState, useEffect } from "react";
import { Header } from "../Header/Header";
import AddTodo from "./AddTodo";
import { useQuery } from "@tanstack/react-query";
import api from "../api";

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
      </div>
    </>
  );
};

export default TodoPage;
