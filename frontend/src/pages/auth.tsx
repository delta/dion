import { Route, Routes } from "react-router-dom";

import CallBack from "../components/auth/callback";
import Login from "../components/auth/login";

function AuthPage() {
  return (
    <Routes>
      <Route path="/callback" element={<CallBack />} />
      <Route path="/login" element={<Login />} />
    </Routes>
  );
}

export default AuthPage;
