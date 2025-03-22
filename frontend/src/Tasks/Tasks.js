import React, { useState } from "react";
import { LogOut } from "../modals/LogOut";

const TodoPage = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <>
      <div className="h-full w-full">
        <div
          className="flex justify-end pr-2 pt-2 text-primary cursor-pointer"
          onClick={() => setIsModalOpen(true)}
        >
          Log out
        </div>
      </div>
      <LogOut isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} />
    </>
  );
};

export default TodoPage;
