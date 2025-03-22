import React from "react";
import { useNavigate } from "react-router-dom";

const TodoPage = () => {
  const navigate = useNavigate();

  return (
    <div className="h-full w-full">
      <div
        className="flex justify-end pr-2 pt-2 text-primary"
        onClick={() => {
          navigate("/login");
          localStorage.removeItem("token");
        }}
      >
        Log out
      </div>
    </div>
  );
};

export default TodoPage;
