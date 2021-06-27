import logo from "./logo.svg";
import "./App.css";
import Login from "./pages/Login";
import Nav from "./components/Nav";
import Register from "./pages/Register";
import Expenses from "./pages/Expenses";
import { BrowserRouter, Route } from "react-router-dom";
import React, { useState, useEffect } from "react";
import "./config";

function App() {
  const [firstName, setFirstName] = useState("");
  const [loginPageMsg, setLoginPageMsg] = useState("");

  useEffect(() => {
    (async () => {
      const response = await fetch(`${global.config.baseURL}/api/user`, {
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      const content = await response.json();
      setFirstName(content.firstName);
    })();
  });

  return (
    <div className="App">
      <BrowserRouter>
        <Nav
          firstName={firstName}
          setFirstName={setFirstName}
          setLoginPageMsg={setLoginPageMsg}
        />
        <Route
          path="/"
          exact
          component={() => (
            <Login
              setFirstName={setFirstName}
              loginPageMsg={loginPageMsg}
              setLoginPageMsg={setLoginPageMsg}
            />
          )}
        />
        <Route
          path="/login"
          component={() => (
            <Login
              setFirstName={setFirstName}
              loginPageMsg={loginPageMsg}
              setLoginPageMsg={setLoginPageMsg}
            />
          )}
        />
        <Route
          path="/register"
          component={() => <Register setLoginPageMsg={setLoginPageMsg} />}
        />
        <Route
          path="/expenses"
          component={() => (
            <Expenses firstName={firstName} setFirstName={setFirstName} />
          )}
        />
      </BrowserRouter>
    </div>
  );
}

export default App;
