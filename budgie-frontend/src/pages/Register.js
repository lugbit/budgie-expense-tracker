import React, { useState } from "react";
import { Redirect } from "react-router";
import { Link } from "react-router-dom";
import "../config";

// Register page
const Register = (props) => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passwordConfirm, setPasswordConfirm] = useState("");
  const [redirect, setRedirect] = useState(false);
  const [errors, setErrors] = useState([]);

  const submit = async (e) => {
    e.preventDefault();
    setErrors([]);

    // check passwords match
    if (password == passwordConfirm) {
      const response = await fetch(`${global.config.baseURL}/api/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          firstName,
          lastName,
          email,
          password,
        }),
      });

      const content = await response.json();

      // check for errors in the response
      if (content.errors != null && content.errors.length != 0) {
        for (var i = 0; i < content.errors.length; i++) {
          setErrors((errors) => [...errors, content.errors[i].message]);
        }
      } else {
        // if no errors, redirect
        setRedirect(true);
        props.setLoginPageMsg("Registration Successful! Please login");
      }
    } else {
      alert("Passwords do not match");
    }
  };

  if (redirect) {
    return <Redirect to="/login/" />;
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
      <main className="form-signin">
        <form onSubmit={submit}>
          <h1 className="h3 mb-3 fw-normal">Register</h1>
          <input
            type="text"
            className="form-control"
            placeholder="First name"
            onChange={(e) => setFirstName(e.target.value)}
            required
          />

          <input
            type="text"
            className="form-control"
            placeholder="Last name"
            onChange={(e) => setLastName(e.target.value)}
            required
          />

          <input
            type="email"
            className="form-control"
            placeholder="Email address"
            onChange={(e) => setEmail(e.target.value)}
            required
          />

          <input
            type="password"
            className="form-control"
            placeholder="Password"
            onChange={(e) => setPassword(e.target.value)}
            required
          />

          <input
            type="password"
            className="form-control"
            placeholder="Confirm Password"
            onChange={(e) => setPasswordConfirm(e.target.value)}
            required
          />

          <button className="w-100 btn btn-lg btn-dark" type="submit">
            Submit
          </button>
          <p>
            Already have an account? <Link to="/login">Click here</Link> to
            login.
          </p>
        </form>
      </main>
    </div>
  );
};

export default Register;
