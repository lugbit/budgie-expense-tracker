import React from "react";
import { Link } from "react-router-dom";
import logo from "../bird_logo.svg";
import "../config";

const Nav = (props) => {
  const logout = async () => {
    await fetch(`${global.config.baseURL}/api/logout`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    props.setFirstName("");
    props.setLoginPageMsg("Logout Successful!");
  };

  let menu;

  if (props.firstName === "" || props.firstName == null) {
    menu = (
      <ul class="navbar-nav mr-auto">
        <li class="nav-item active">
          <Link to="/login" className="nav-link active" aria-current="page">
            Login
          </Link>
        </li>
        <li class="nav-item active">
          <Link to="/register" className="nav-link active" aria-current="page">
            Register
          </Link>
        </li>
      </ul>
    );
  } else {
    menu = (
      <ul className="navbar-nav me-auto mb-2 mb-md-0">
        <li className="nav-item">
          <Link
            to="/login"
            className="nav-link active"
            aria-current="page"
            onClick={logout}
          >
            Logout
          </Link>
        </li>
      </ul>
    );
  }

  return (
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
      <a class="navbar-brand">
        <img
          src={logo}
          width="30"
          height="30"
          class="d-inline-block align-top"
          alt="Budgie Budget Bird"
        ></img>
        &nbsp; Budgie Expense Tracker
      </a>
      <button
        class="navbar-toggler"
        type="button"
        data-toggle="collapse"
        data-target="#navbarsExampleDefault"
        aria-controls="navbarsExampleDefault"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        {menu}
      </div>
    </nav>
  );
};

export default Nav;
