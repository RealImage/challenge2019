const fs = require("fs");
const path = require("path");

module.exports.writeFirstResult = dataList => {
  try {
    dataList.forEach(dataRow => {
      fs.appendFileSync(
        path.resolve(__dirname, "../output", "out1.csv"),
        dataRow + "\n",
        "utf8",
        err => {
          if (err) throw err;
        }
      );
    });
  } catch (e) {
    throw e;
  }
};

module.exports.writeSecondResult = dataList => {
  try {
    dataList.forEach(dataRow => {
      fs.appendFileSync(
        path.resolve(__dirname, "../output", "out2.csv"),
        dataRow + "\n",
        "utf8",
        err => {
          if (err) throw err;
        }
      );
    });
  } catch (e) {
    throw e;
  }
};
