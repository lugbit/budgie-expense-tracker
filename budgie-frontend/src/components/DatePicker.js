import React from "react";
import DateRangePicker from "@wojtekmaj/react-daterange-picker";
import { IoCalendar } from "react-icons/io5";

const DatePicker = ({ dateRangeValue, setDateRangeValue, addDays }) => {
  return (
    <div>
      <form class=" date-range-form form-inline">
        <div class="form-group">
          <label className="date-range-form-labels" for="date-range">
            Date Range:
          </label>
          <DateRangePicker
            onChange={setDateRangeValue}
            value={dateRangeValue}
            required
            calendarIcon={<IoCalendar />}
            rangeDivider="to"
          />
        </div>
        <div class="form-group">
          <div
            className="date-range-form-labels"
            role="button"
            onClick={() =>
              setDateRangeValue([addDays(new Date(), -7), new Date()])
            }
          >
            Last 7 Days
          </div>
        </div>
        |
        <div class="form-group">
          <div
            className="date-range-form-labels"
            role="button"
            onClick={() =>
              setDateRangeValue([addDays(new Date(), -30), new Date()])
            }
          >
            Last 30 Days
          </div>
        </div>
      </form>
    </div>
  );
};

export default DatePicker;
