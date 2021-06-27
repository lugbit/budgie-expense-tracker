import React, { useEffect } from "react";
import { PieChart } from "react-minimal-pie-chart";

const ExpenseCatCountGraph = ({ data }) => {
  return (
    <div>
      {data.length > 0 ? (
        <div>
          <div className="expenseForm shadow p-3 mb-5 bg-white rounded">
            <h3 className="chart-header">Category Breakdown</h3>
            <div className="chart-container">
              <PieChart
                animate
                animationDuration={500}
                animationEasing="ease-out"
                center={[50, 50]}
                data={data}
                lengthAngle={360}
                lineWidth={15}
                paddingAngle={0}
                radius={50}
                rounded
                startAngle={0}
                viewBoxSize={[100, 100]}
                label={(data) => data.dataEntry.title}
                labelPosition={65}
                labelStyle={{
                  fontSize: "6px",
                  fontColor: "FFFFFA",
                  fontWeight: "400",
                }}
              />
            </div>
          </div>
        </div>
      ) : (
        <div></div>
      )}
    </div>
  );
};

export default ExpenseCatCountGraph;
