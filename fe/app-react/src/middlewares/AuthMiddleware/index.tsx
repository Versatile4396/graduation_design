import { Outlet, useNavigate } from "react-router";
import { useEffect, useState } from "react";
import { Spin } from "antd";

const AuthMiddleware = () => {
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  useEffect(() => {
    navigate("/login");
    console.log("AuthMiddleware 中间件触发");
  }, []);
  return !loading ? <Outlet /> : <Spin spinning />;
};

export default AuthMiddleware;
