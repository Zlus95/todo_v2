import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Auth from "../Auth/Auth";
import TodoPage from "../Tasks/Tasks";

const Navigation = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<TodoPage />} />
        <Route path="/login" element={<Auth />} />
      </Routes>
    </Router>
  );
};

export default Navigation;
