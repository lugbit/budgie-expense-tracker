import React, { useEffect } from "react";

// Alert component with timeout
const Alert = ({ type, msg, removeAlert }) => {
  useEffect(() => {
    const timeout = setTimeout(() => {
      removeAlert();
    }, 3000);
    return () => clearTimeout(timeout);
  }, []);
  return (
    <div className={`alert alert-${type}`} role={alert}>
      {msg}
    </div>
  );
};

export default Alert;
