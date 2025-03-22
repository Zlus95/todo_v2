import React, { memo, useCallback } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import api from "../api";

async function updateStatus(id, todo) {
  const task = await todo.find((item) => item.id === id);
  const { data } = await api.patch(`/tasks/${id}`, {
    ...task,
    done: !task.done,
  });
  return data;
}

async function deleteTask(id) {
  const { data } = await api.delete(`/tasks/${id}`);
  return data;
}

const Todo = (props) => {
  const queryClient = useQueryClient();
  const { title, id, todo, status } = props;
  
  const mutationUpdate = useMutation({
    mutationFn: (id) => updateStatus(id, todo.data),
    onSuccess: () => queryClient.invalidateQueries(["todoList"]),
  });

  const update = useCallback(async () => {
    try {
      await mutationUpdate.mutateAsync(id);
    } catch (error) {
      console.error("error", error);
      alert("Failed to update task status. Please try again");
    }
  }, [id, mutationUpdate]);

  const mutationDelete = useMutation({
    mutationFn: (id) => deleteTask(id),
    onSuccess: () => queryClient.invalidateQueries(["todoList"]),
  });

  const deleteCallBack = useCallback(async () => {
    try {
      await mutationDelete.mutateAsync(id);
    } catch (error) {
      console.error("error", error);
      alert("Failed to delete task Please try again");
    }
  }, [id, mutationDelete]);

  return (
    <>
      {title}
      <div>
        <button onClick={update}>{status}</button>
        <button onClick={() => deleteCallBack(id)}>x</button>
      </div>
    </>
  );
};

export default memo(Todo);
