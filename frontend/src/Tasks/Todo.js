import React, { memo, useCallback, useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import api from "../api";
import { Edit } from "../modals/Edit";

async function updateStatus(id, title, status) {
  const { data } = await api.patch(`/task/${id}`, {
    title: title,
    status: status,
  });
  return data;
}

async function deleteTask(id) {
  const { data } = await api.delete(`/task/${id}`);
  return data;
}

const Todo = (props) => {
  const queryClient = useQueryClient();
  const { title, id, status } = props;
  const [editModal, setEditModal] = useState(false);

  const mutationUpdate = useMutation({
    mutationFn: ({ id, title, status }) => updateStatus(id, title, status),
    onSuccess: () => queryClient.invalidateQueries(["todoList"]),
  });

  const update = useCallback(
    async (id, title, status) => {
      try {
        await mutationUpdate.mutateAsync({ id, title, status });
      } catch (error) {
        console.error("error", error);
        alert("Failed to update task status. Please try again");
      }
    },
    [mutationUpdate]
  );

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
      <div className="flex gap-2">
        <button onClick={() => setEditModal(true)}>{status}</button>
        <button className="text-red-500" onClick={() => deleteCallBack(id)}>
          x
        </button>
      </div>
      <Edit
        isOpen={editModal}
        onClose={() => setEditModal(false)}
        update={update}
        status={status}
        id={id}
        title={title}
      />
    </>
  );
};

export default memo(Todo);
