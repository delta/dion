import { BrowserRouter, Route, Routes } from "react-router-dom";

import { AuthContextProvider } from "./context/user-context";
import AuthPage from "./pages/auth";

function App() {
  // Dummy code that works
  // useEffect(() => {
  //   init().then(() => {
  //     const a = encrypt(JSON.stringify({"a": "b"}), new_key());
  //     console.log(a.get_nonce());
  //     console.log(a.get_value());
  //   })
  // })

  return (
    <AuthContextProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/auth/*" element={<AuthPage />} />
        </Routes>
      </BrowserRouter>
    </AuthContextProvider>
  );
}

export default App;
