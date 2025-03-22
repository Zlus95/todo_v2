import React, { useState } from "react";

const options = ["to do", "doing", "done"];

export const Edit = ({ isOpen, onClose, status, update, title, id }) => {
  const [value, setValue] = useState(title || "");
  const [isOpenDrop, setIsOpenOpenDrop] = useState(false);
  const [selectedOption, setSelectedOption] = useState(status);

  const handleSelect = (option) => {
    setSelectedOption(option);
    setIsOpenOpenDrop(false);
  };

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
        <input
          type="text"
          onChange={(e) => setValue(e.target.value)}
          value={value}
          id="todo"
        />
        <div className="relative mt-2">
          <button
            onClick={() => setIsOpenOpenDrop(!isOpenDrop)}
            className="bg-white border border-gray-300 rounded-md px-4 py-2 flex items-center justify-between w-48"
          >
            <span>{selectedOption || "Выберите опцию"}</span>
            <svg
              className={`w-5 h-5 ml-2 transition-transform ${
                isOpenDrop ? "rotate-180" : ""
              }`}
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M19 9l-7 7-7-7"
              />
            </svg>
          </button>

          {isOpenDrop && (
            <div className="absolute mt-2 w-48 bg-white border border-gray-300 rounded-md shadow-lg">
              {options.map((option, index) => (
                <div
                  key={index}
                  onClick={() => handleSelect(option)}
                  className="px-4 py-2 hover:bg-gray-100 cursor-pointer"
                >
                  {option}
                </div>
              ))}
            </div>
          )}
        </div>
        <div className="flex justify-between">
          <button
            onClick={onClose}
            className="mt-4 bg-zinc-600 text-white px-4 py-2 rounded"
          >
            Close
          </button>
          <button
            onClick={() => {
              update(id, value, selectedOption);
              onClose();
            }}
            className={
              value.length > 0
                ? "mt-4 bg-primary text-white px-4 py-2 rounded"
                : "mt-4 bg-primary text-white px-4 py-2 rounded opacity-50"
            }
            disabled={!value.length}
          >
            Edit
          </button>
        </div>
      </div>
    </div>
  );
};
