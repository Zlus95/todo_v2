import React from "react";
import { useNavigate } from "react-router-dom";

export const Delete = ({ isOpen, onClose, id, deleteCallBack }) => {
  const navigate = useNavigate();

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white p-6 rounded-lg shadow-lg relative">
        <button
          onClick={onClose}
          className="absolute top-2 right-2 text-gray-500 hover:text-gray-700"
        >
          &times;
        </button>
        <div className="text-lg">
          Are you sure you want to delete this task?
        </div>
        <div className="flex justify-between">
          <button
            onClick={onClose}
            className="mt-4 bg-zinc-600 text-white px-4 py-2 rounded"
          >
            Close
          </button>
          <button
            onClick={() => deleteCallBack(id)}
            className="mt-4 bg-red-500 text-white px-4 py-2 rounded"
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  );
};
