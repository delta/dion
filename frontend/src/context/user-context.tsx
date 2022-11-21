import React, { createContext, useEffect, useState } from "react";

interface AuthContextType {
  email: string | null;
  loading: boolean;
  loginUser: (email: string) => void;
  logoutUser: () => void;
}

export const AuthContext = createContext<AuthContextType | null>(null);

interface AuthUserJson {
  email: string;
}

interface Props {
  children: React.ReactNode;
}

export const AuthContextProvider = ({ children }: Props) => {
  const [loading, setLoading] = useState(true);
  const [email, setEmail] = useState<string | null>(null);
  const loginUser = (email: string) => {
    setEmail(email);
  };
  const logoutUser = () => {
    setEmail(null);
  };

  useEffect(() => {
    void (async () => {
      if (loading) {
        const resp = await fetch("http://localhost:8000/auth/user", {
          credentials: "include",
        });
        if (resp.status === 200) {
          try {
            const json: AuthUserJson = await resp.json();
            loginUser(json.email);
          } catch (e) {
            console.log(e);
          }
        }
      }
      setLoading(false);
    })();
  }, []);
  if (loading) {
    return <h1>Loading</h1>;
  }
  return (
    <AuthContext.Provider value={{ email, loading, loginUser, logoutUser }}>
      {children}
    </AuthContext.Provider>
  );
};
