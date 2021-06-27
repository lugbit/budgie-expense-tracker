import React, { useState } from "react";
import { Redirect } from "react-router";
import { Link } from "react-router-dom";
import "../config";

// Login page
const Login = (props) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false);
  const [errors, setErrors] = useState([]);

  const submit = async (e) => {
    e.preventDefault();
    setErrors([]);

    const response = await fetch(`${global.config.baseURL}/api/authenticate`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        email,
        password,
      }),
    });

    const content = await response.json();

    if (content.errors != null && content.errors.length != 0) {
      for (var i = 0; i < content.errors.length; i++) {
        setErrors((errors) => [...errors, content.errors[i].message]);
      }
    } else {
      setRedirect(true);
    }

    props.setFirstName(content.firstName);
  };

  if (redirect) {
    return <Redirect to="/expenses" />;
  }

  return (
    <div>
      {errors.map((error, index) => (
        <div
          key={index}
          className="alert alert-danger other-alert"
          role="alert"
        >
          {error}
        </div>
      ))}
      {props.loginPageMsg && (
        <div className="alert alert-success other-alert" role="alert">
          {props.loginPageMsg}
        </div>
      )}
      <main className="form-signin">
        <form onSubmit={submit}>
          <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

          <input
            type="email"
            className="form-control"
            placeholder="Email"
            required
            onChange={(e) => setEmail(e.target.value)}
          />

          <input
            type="password"
            className="form-control"
            placeholder="Password"
            required
            onChange={(e) => setPassword(e.target.value)}
          />

          <button className="w-100 btn btn-lg btn-dark" type="submit">
            Sign in
          </button>
          <p>
            Don't have an account? <Link to="/register">Click here</Link> to
            register.
          </p>
        </form>
      </main>
    </div>
  );
};

export default Login;
