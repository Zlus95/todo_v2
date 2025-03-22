import React, { useState } from "react";
import { LogOut } from "../modals/LogOut";

export const Header = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);

  return (
    <>
      <div
        className="flex justify-end pr-2 pt-2 text-primary cursor-pointer"
        onClick={() => setIsModalOpen(true)}
      >
        Log out
      </div>
      <LogOut isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} />
    </>
  );
};
