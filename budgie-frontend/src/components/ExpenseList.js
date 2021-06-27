import React from "react";

import {
  IoTrash,
  IoCreate,
  IoPricetags,
  IoRestaurant,
  IoBuild,
  IoHome,
  IoBaseball,
  IoStar,
  IoCalendar,
  IoChatbox,
} from "react-icons/io5";

import { GrNext, GrPrevious } from "react-icons/gr";

// currency formatter to format numbers to AUD
let currencyFormatter = new Intl.NumberFormat(undefined, {
  style: "currency",
  currency: "AUD",

  // These options are needed to round to whole numbers if that's what you want.
  //minimumFractionDigits: 0, // (this suffices for whole numbers, but will print 2500.10 as $2,500.1)
  //maximumFractionDigits: 0, // (causes 2500.99 to be printed as $2,501)
});

// ExpenseList displays the list of expenses
const ExpenseList = ({
  items,
  count,
  total,
  offset,
  removeItem,
  editItem,
  paginationNext,
  paginationPrev,
}) => {
  let displayList;

  // function to return an icon based on the category
  const showCatIcon = (categoryName) => {
    let icon;
    if (categoryName == "Uncategorized") {
      icon = <IoPricetags />;
    } else if (categoryName == "Food") {
      icon = <IoRestaurant />;
    } else if (categoryName == "Utility") {
      icon = <IoBuild />;
    } else if (categoryName == "Rent/Mortgage") {
      icon = <IoHome />;
    } else if (categoryName == "Hobbies") {
      icon = <IoBaseball />;
    } else {
      icon = <IoStar />;
    }
    return icon;
  };

  // if the list is not empty, assign these values into displayList
  if (items != null && items.length > 0) {
    displayList = (
      <div>
        <table class="table">
          <tbody>
            {items.map((item) => {
              const { expenseID, title, categoryName, date, amount, note } =
                item;

              return (
                <tr>
                  <th scope="row">
                    <p className="lead">
                      {showCatIcon(categoryName)}{" "}
                      {currencyFormatter.format(amount)}
                    </p>
                  </th>
                  <td>
                    <p className="lead">
                      {title.length > 25 ? title.slice(0, 25) + "..." : title}{" "}
                      {note && <IoChatbox />}
                    </p>
                  </td>
                  <td>
                    <p className="lead">
                      {date} <IoCalendar />
                    </p>
                  </td>
                  <td>
                    <div type="button" onClick={() => editItem(expenseID)}>
                      <p className="lead">
                        <IoCreate />
                      </p>
                    </div>
                  </td>
                  <td>
                    <div type="button" onClick={() => removeItem(expenseID)}>
                      <p className="lead">
                        <IoTrash />
                      </p>
                    </div>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    );
  } else {
    displayList = <p className="no-expense-title">No expenses found</p>;
  }

  return (
    <div>
      {count > 0 && (
        <div className="row">
          <div className="col-sm-6">
            <p className="expense-list-result-header">
              {count} result(s) found
            </p>
          </div>
          <div className="col-sm-6">
            <h3 className="expense-total">
              Total: {currencyFormatter.format(total)}
            </h3>
          </div>
        </div>
      )}

      {displayList}

      {offset != 0 && (
        <GrPrevious size={30} type="button" onClick={() => paginationPrev()} />
      )}

      {offset + 5 < count && (
        <GrNext size={30} type="button" onClick={() => paginationNext()} />
      )}
    </div>
  );
};

export default ExpenseList;
