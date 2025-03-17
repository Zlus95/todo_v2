import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Auth from "../Auth/Auth";
import TodoPage from "../Tasks/Tasks";
import Registration from "../Auth/Registration";

const withAuth = (Component) => {
  return (props) => {
    const navigate = useNavigate();
    const token = localStorage.getItem("token");

    useEffect(() => {
      if (!token) {
        navigate("/login");
      }
    }, [token, navigate]);

    if (!token) {
      return null;
    }
    return <Component {...props} />;
  };
};

const PrivateTodoPage = withAuth(TodoPage);

const Navigation = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<PrivateTodoPage />} />
        <Route path="/login" element={<Auth />} />
        <Route path="/register" element={<Registration />} />
      </Routes>
    </Router>
  );
};

export default Navigation;
