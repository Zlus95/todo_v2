import React, { memo, useCallback, useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import api from "../api";
import { Edit } from "../modals/Edit";
import { Delete } from "../modals/Delete";

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

const styleStatus = {
  "to do": "text-primary",
  doing: "text-green-500",
  done: "text-stone-500",
};

const Todo = (props) => {
  const queryClient = useQueryClient();
  const { title, id, status } = props;
  const [editModal, setEditModal] = useState(false);
  const [deletetModal, setDeleteModal] = useState(false);

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
    mutationFn: ({ id }) => deleteTask(id),
    onSuccess: () => queryClient.invalidateQueries(["todoList"]),
  });

  const deleteCallBack = useCallback(
    async (id) => {
      try {
        await mutationDelete.mutateAsync({ id });
      } catch (error) {
        console.error("error", error);
        alert("Failed to delete task Please try again");
      }
    },
    [mutationDelete]
  );

  return (
    <>
      {title}
      <div className="flex gap-2">
        <button
          onClick={() => setEditModal(true)}
          className={styleStatus[status]}
        >
          {status}
        </button>
        <button className="text-red-500" onClick={() => setDeleteModal(true)}>
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
      <Delete
        isOpen={deletetModal}
        onClose={() => setDeleteModal(false)}
        id={id}
        deleteCallBack={deleteCallBack}
      />
    </>
  );
};

export default memo(Todo);
