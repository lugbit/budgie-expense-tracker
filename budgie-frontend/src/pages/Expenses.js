import React, { useState, useEffect } from "react";
import { withRouter, Link } from "react-router-dom";
import Alert from "../components/Alert";
import ExpenseList from "../components/ExpenseList";
import ExpenseCatCountGraph from "../components/ExpenseCatCountGraph";
import ExpenseForm from "../components/ExpenseForm";
import "../config";
import DatePicker from "../components/DatePicker";

// addDays takes a date and adds number of days specified
let addDays = (originalDate, days) => {
  // clone input date
  let cloneDate = new Date(originalDate.valueOf());
  // add number of days
  cloneDate.setDate(cloneDate.getDate() + days);
  // return new date
  return cloneDate;
};

// default date is today
let today = new Date();
let dd = String(today.getDate()).padStart(2, "0");
let mm = String(today.getMonth() + 1).padStart(2, "0"); //January is 0!
let yyyy = today.getFullYear();

let todaysDate = yyyy + "-" + mm + "-" + dd;

// Expenses page
const Expenses = ({ firstName }) => {
  // form properties
  const [title, setTitle] = useState("");
  const [expenseCategoryID, setExpenseCategoryID] = useState("1");
  const [date, setDate] = useState(todaysDate);
  const [amount, setAmount] = useState("");
  const [note, setNote] = useState("");
  const [submittedFlag, setSubmittedFlag] = useState(false);
  const [removedFlag, setRemovedFlag] = useState(false);

  // edit flag and edit ID
  const [isEditing, setIsEditing] = useState(false); // true when in edit mode
  const [editID, setEditID] = useState(null); // id of the expense currently being edited

  // expense list properties
  const [expenseList, setExpenseList] = useState([]); // expense array
  const [expenseListCount, setExpenseListCount] = useState(0); // count of expenses
  const [expenseListTotal, setExpenseListTotal] = useState(0); // expenses total
  const [offset, setOffset] = useState(0); // pagination offset

  // expense category array (categories are pulled from db)
  const [categories, setCategories] = useState([]);

  // datetime range picker value. Defaults to last 7 days
  const [dateRangeValue, setDateRangeValue] = useState([
    addDays(new Date(), -7),
    new Date(),
  ]);

  // alert usestate
  const [alert, setAlert] = useState({
    show: false,
    msg: "",
    type: "",
  });

  // pie chart
  const [pieData, setPieData] = useState([]);

  // pie chart category constants
  // https://medialab.github.io/iwanthue/
  const COLOR_UNCATEGORIZED = "#85fff6";
  const COLOR_FOOD = "#ffbd8b";
  const COLOR_UTILITY = "#37bef0";
  const COLOR_RENT_MORTGAGE = "#8dbb79";
  const COLOR_HOBBIES = "#b1aff8";
  const COLOR_OTHER = "#ffc9ec";

  // useEffect to fetch expenses. Includes a date range from the datetime range picker
  useEffect(() => {
    (async () => {
      // only perform fetch if dateRangeValue is not null
      if (dateRangeValue != null) {
        const response = await fetch(`${global.config.baseURL}/api/expenses`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          credentials: "include", // send cookie to authorize
          body: JSON.stringify({
            startDate: dateRangeValue[0], // start date of the date picker
            endDate: dateRangeValue[1], // end date of the date picker
            limit: 5,
            offset: offset,
          }),
        });

        const content = await response.json();
        // set expenseList
        setExpenseListCount(content.count);
        setExpenseListTotal(content.total);
        setExpenseList(content.expenses);
      }
    })();
  }, [dateRangeValue, submittedFlag, removedFlag, offset]); //only update when

  // useEffect to fetch expense categories
  useEffect(() => {
    (async () => {
      const response = await fetch(
        `${global.config.baseURL}/api/expense-categories`,
        {
          headers: { "Content-Type": "application/json" },
          credentials: "include",
        }
      );

      const content = await response.json();
      // after fetching categories, set the useState
      setCategories(content);
    })();
  }, []);

  // useEffect to fetch expense category count
  useEffect(() => {
    (async () => {
      // only perform fetch if dateRangeValue is not null
      if (dateRangeValue != null) {
        const response = await fetch(
          `${global.config.baseURL}/api/expenses/count-by-category`,
          {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include", // send cookie to authorize
            body: JSON.stringify({
              startDate: dateRangeValue[0], // start date of the date picker
              endDate: dateRangeValue[1], // end date of the date picker
            }),
          }
        );

        const content = await response.json();
        let data = [];
        let categoryColor;

        for (var i = 0; i < content.length; i++) {
          switch (content[i].categoryName) {
            case "Uncategorized":
              categoryColor = COLOR_UNCATEGORIZED;
              break;
            case "Food":
              categoryColor = COLOR_FOOD;
              break;
            case "Utility":
              categoryColor = COLOR_UTILITY;
              break;
            case "Rent/Mortgage":
              categoryColor = COLOR_RENT_MORTGAGE;
              break;
            case "Hobbies":
              categoryColor = COLOR_HOBBIES;
              break;
            case "Other":
              categoryColor = COLOR_OTHER;
              break;
          }
          let insert = {
            color: categoryColor,
            title: content[i].categoryName,
            value: parseInt(content[i].categoryCount),
          };

          data.push(insert);
        }
        setPieData(data);
      }
    })();
  }, [dateRangeValue, submittedFlag, removedFlag]);

  // submit or update expense
  const submitNewExpense = async (e) => {
    // prefent page refresh
    e.preventDefault();
    setSubmittedFlag(true);

    // check that required field is not empty
    if (!title || !amount) {
      // display alert
      showAlert(true, "Please fill out required fields", "danger");
    } else if (isEditing) {
      // edit mode
      const response = await fetch(
        `${global.config.baseURL}/api/expense/${editID}`,
        {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          credentials: "include",
          body: JSON.stringify({
            title,
            categoryID: expenseCategoryID,
            date,
            amount,
            note,
          }),
        }
      );

      const content = await response.json();
      // reset form and show alert
      resetInputs();
      showAlert(true, "Updated successfully!", "success");
    } else {
      //submit new expense
      const response = await fetch(`${global.config.baseURL}/api/expense`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          title,
          categoryID: expenseCategoryID,
          date,
          amount,
          note,
        }),
      });
      const content = await response.json();
      resetInputs();
      showAlert(true, "Added successfully!", "success");
    }

    setSubmittedFlag(false);
  };

  // removeItem accepts an expense ID and makes an API call to delete that item
  const removeItem = async (id) => {
    // confirmation alert
    const r = window.confirm("Are you sure you want to delete this item?");
    if (r == true) {
      setRemovedFlag(true);
      await fetch(`${global.config.baseURL}/api/expense/${id}`, {
        method: "DELETE",
        credentials: "include",
      });

      showAlert(true, "Item deleted successfully!", "success");
      setRemovedFlag(false);
    }
  };

  // For pagination. Add or subtract from the offset
  const paginationPrev = () => {
    setOffset(offset - 5);
  };

  const paginationNext = () => {
    setOffset(offset + 5);
  };

  // show alert helper function
  const showAlert = (show = false, msg = "", type = "") => {
    setAlert({ show, msg, type });
  };

  // editItem accepts an expenseID and fetches that expenses properties
  const editItem = async (id) => {
    const response = await fetch(`${global.config.baseURL}/api/expense/${id}`, {
      method: "GET",
      credentials: "include",
    });
    const content = await response.json();
    // set form input fields with the expense item to be edited
    setTitle(content.title);
    setExpenseCategoryID(content.categoryID);
    setDate(content.date);
    setAmount(content.amount);
    if (!content.note) {
      setNote("");
    } else {
      setNote(content.note);
    }
    // set editing flag to true and edit ID
    setIsEditing(true);
    setEditID(content.expenseID);
  };

  // resets expense form
  const resetInputs = () => {
    setTitle("");
    setExpenseCategoryID("1");
    setDate(todaysDate);
    setAmount("");
    setNote("");
    setEditID(null);
    setIsEditing(false);
  };

  const resetForm = () => {
    resetInputs();
    showAlert(true, "Form was reset", "success");
  };

  let render;
  // if firstName is undefined, show unauthorized page as user is not logged in.
  if (typeof firstName === "undefined") {
    render = (
      <div>
        <h3>Unauthorized</h3>
        <Link to="/login">Click here to login</Link>
      </div>
    );
  } else {
    render = (
      <div>
        <div className="row">
          <div className="col-sm-4">
            <ExpenseCatCountGraph data={pieData} />

            {alert.show && (
              <div className="expense-form-alert">
                <Alert {...alert} removeAlert={showAlert} />
              </div>
            )}

            <ExpenseForm
              submitNewExpense={submitNewExpense}
              isEditing={isEditing}
              title={title}
              setTitle={setTitle}
              expenseCategoryID={expenseCategoryID}
              setExpenseCategoryID={setExpenseCategoryID}
              categories={categories}
              date={date}
              setDate={setDate}
              amount={amount}
              setAmount={setAmount}
              note={note}
              setNote={setNote}
              resetForm={resetForm}
            />
          </div>
          <div className="col-sm-8">
            <DatePicker
              dateRangeValue={dateRangeValue}
              setDateRangeValue={setDateRangeValue}
              addDays={addDays}
            />

            <div className="expenseList shadow-none p-3 mb-5 bg-light rounded">
              <ExpenseList
                items={expenseList}
                count={expenseListCount}
                total={expenseListTotal}
                offset={offset}
                removeItem={removeItem}
                editItem={editItem}
                paginationNext={paginationNext}
                paginationPrev={paginationPrev}
              />
            </div>
            <div></div>
          </div>
        </div>
      </div>
    );
  }
  return render;
};

export default withRouter(Expenses);
