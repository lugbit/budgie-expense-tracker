import React from "react";




const ExpenseForm = ({
  submitNewExpense,
  isEditing,
  title,
  setTitle,
  expenseCategoryID,
  setExpenseCategoryID,
  categories,
  date,
  setDate,
  amount,
  setAmount,
  note,
  setNote,
  resetForm,
}) => {
  return (
    <div className="expenseForm shadow p-3 mb-5 bg-white rounded">
      <form onSubmit={submitNewExpense}>
        <div className="form-group">
          <h3 className="expense-headers">
            {isEditing ? "Update Expense" : "Add Expense"}
          </h3>
        </div>

        <div class="form-row">
          <div class="col">
            <div className="form-group">
              <div class="input-group">
                <input
                  type="text"
                  className="form-control"
                  placeholder="Expense title"
                  maxLength="40"
                  required
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                />
              </div>
            </div>
          </div>
          <div class="col">
            <div className="form-group">
              <select
                className="form-control"
                value={expenseCategoryID}
                onChange={(e) => setExpenseCategoryID(e.target.value)}
              >
                {categories.map((category, index) => (
                  <option value={category.categoryID}>
                    {category.categoryName}
                  </option>
                ))}
              </select>
            </div>
          </div>
        </div>
        <div class="form-row">
          <div class="col">
            <div className="form-group">
              <input
                type="date"
                className="form-control"
                placeholder="Date"
                value={date}
                onChange={(e) => setDate(e.target.value)}
              />
            </div>
          </div>
          <div class="col">
            <div className="form-group">
              <div class="input-group mb-3">
                <div class="input-group-prepend">
                  <span class="input-group-text">$</span>
                </div>
                <input
                  type="number"
                  min="0.1"
                  max="1000000"
                  step="any"
                  className="form-control"
                  placeholder="Amount"
                  required
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                />
                
              </div>
            </div>
          </div>
        </div>
        <div class="form-row">
          <div class="col">
            <div className="form-group">
              <input
                type="text"
                className="form-control"
                placeholder='Add a note e.g. "Food for Angie the budgie"'
                maxLength="100"
                value={note}
                onChange={(e) => setNote(e.target.value)}
              />
            </div>
          </div>
        </div>

        <div className="form-group">
          <div className="row">
            <div className="col-sm-3">
              <button
                className="w-100 btn btn-outline-info"
                onClick={() => resetForm()}
              >
                Reset
              </button>
            </div>
            <div className="col-sm-9">
              {isEditing ? (
                <button className="w-100 btn btn-outline-warning" type="submit">
                  Update
                </button>
              ) : (
                <button className="w-100 btn btn-outline-success" type="submit">
                  Add
                </button>
              )}
            </div>
          </div>
        </div>
      </form>
    </div>
  );
};

export default ExpenseForm;
