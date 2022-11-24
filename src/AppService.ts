import fs from "fs";
import csv from "csv-parser";
import { csvParseService } from "./CsvParseService";

interface Result {
  distributor: string;
  deliverable: string;
  partners: string;
  amount: string;
}
export const appService = new (class AppService {
  public validateAndSolveProblemOne = async (parseData: any) => {
    const finalArrayList: Result[] = [];
    const results: any[] = [];
    const resultData = fs
      .createReadStream("./input.csv")
      .pipe(csv())
      .on("data", (data) => {
        results.push(data);
      })
      .on("end", async () => {
        /*In input csv no headers mentioned,so that first data will considered as header
          So,I add headers also in array to process
        */
        const finalInputArray = [];
        const getHeader: string[] = Object.keys(results[0]);
        let headerobject = {
          "150": getHeader[0],
          D1: getHeader[1],
          T1: getHeader[2],
        };
        finalInputArray.push(headerobject);
        results.map((data) => finalInputArray.push(data));
        await appService.SolveProblemOne(parseData, finalInputArray);
        await appService.solveProblemTwo(parseData, finalInputArray);
      });
  };
  public SolveProblemOne = async (parseData: any[], result: any[]) => {
    let finalArrayList: Result[] = [];
    result.map((data) => {
      const filteredData = parseData.filter(
        (filterData) => String(filterData["Theatre"].trim()) === data["T1"]
      );
      const resultedFilterSize = filteredData.filter((filterSize) => {
        const sizeRangeArray = filterSize["Size Slab (in GB)"].split("-");
        const minimum_range = Math.min(
          parseInt(sizeRangeArray[0]),
          parseInt(String(sizeRangeArray[1]).trim())
        );
        const maximum_range = Math.max(
          parseInt(sizeRangeArray[0]),
          parseInt(String(sizeRangeArray[1]).trim())
        );
        return (
          parseInt(data["150"]) >= minimum_range &&
          parseInt(data["150"]) <= maximum_range
        );
      });
      let resultMap = {
        distributor: "",
        deliverable: "false",
        partners: "",
        amount: "",
      };
      if (!resultedFilterSize.length) {
        resultMap = {
          ...resultMap,
          distributor: data["D1"],
          partners: "''",
          amount: "''",
        };
      }
      resultedFilterSize.map((filterObject, index) => {
        let amount =
          filterObject["Minimum cost"] >
          parseInt(data["150"]) * parseInt(String(filterObject["Cost Per GB"]))
            ? filterObject["Minimum cost"]
            : parseInt(data["150"]) *
              parseInt(String(filterObject["Cost Per GB"]));
        if (index === 0) {
          resultMap = {
            ...resultMap,
            distributor: data["D1"],
            partners: filterObject["Partner ID"],
            deliverable: "true",
            amount: String(amount),
          };
        } else {
          if (amount === parseInt(resultMap["amount"])) {
            resultMap = {
              ...resultMap,
              partners:
                resultMap["partners"] + "," + filterObject["Partner ID"],
            };
          }
          if (amount < resultMap["amount"]) {
            resultMap = {
              ...resultMap,
              distributor: data["D1"],
              partners: filterObject["Partner ID"],
              deliverable: "true",
              amount: String(amount),
            };
          }
        }
      });
      finalArrayList.push(resultMap);
    });
    csvParseService.writeCsvData("./output.csv", finalArrayList);
    csvParseService.writeCsvData("./output1.csv", finalArrayList);
  };

  public solveProblemTwo = async (partnerData: any[], inputData: any[]) => {
    const results: any[] = [];
    fs.createReadStream("./capacities.csv")
      .pipe(csv({}))
      .on("data", (data) => results.push(data))
      .on("end", async () => {
        await appService.validateAndSolveProblemTwo(
          partnerData,
          inputData,
          results
        );
      });
  };

  public validateAndSolveProblemTwo = async (
    partnerData: any[],
    inputData: any[],
    capacityData: any[]
  ) => {
    let finalArrayList: any[] = [];
    const theatreFilter = inputData.filter((data) => data["T1"] === "T1");
    const patnerTheatreFilter = partnerData.filter(
      (data) => data["Theatre"].trim() === "T1"
    );
    let requiredCapacityForT1 = 0;
    theatreFilter.map((data) => {
      requiredCapacityForT1 += parseInt(data["150"]);
    });
    const resultFilter = theatreFilter.map((data) => {
      const resultedFilterSize = patnerTheatreFilter.filter((filterSize) => {
        const sizeRangeArray = filterSize["Size Slab (in GB)"].split("-");
        const minimum_range = Math.min(
          parseInt(sizeRangeArray[0]),
          parseInt(String(sizeRangeArray[1]).trim())
        );
        const maximum_range = Math.max(
          parseInt(sizeRangeArray[0]),
          parseInt(String(sizeRangeArray[1]).trim())
        );
        return (
          parseInt(data["150"]) >= minimum_range &&
          parseInt(data["150"]) <= maximum_range
        );
      });
      return resultedFilterSize;
    });
    const filterResultPartnerArray = resultFilter[0].concat(resultFilter[1]);
    let amountPartnerDetails: {
      theatre: string;
      deliverable: string;
      partner: string;
      amount: string;
    }[] = [];
    let amountDetails = {
      theatre: "",
      deliverable: "",
      partner: "",
      amount: "",
    };
    resultFilter[0].map((data) => {
      let amount =
        data["Minimum cost"] >
        parseInt(theatreFilter[0]["150"]) *
          parseInt(String(data["Cost Per GB"]))
          ? data["Minimum cost"]
          : parseInt(theatreFilter[0]["150"]) *
            parseInt(String(data["Cost Per GB"]));
      let deliverable = capacityData.filter(
        (data) =>
          parseInt(String(data["Capacity (in GB)"]).trim()) >=
          parseInt(String(theatreFilter[0]["150"]).trim())
      );

      amountDetails = {
        ...amountDetails,
        theatre: data["Theatre"].trim(),
        deliverable: "true",
        partner: data["Partner ID"].trim(),
        amount: String(amount).trim(),
      };
      amountPartnerDetails.push(amountDetails);
    });
    finalArrayList.push(amountPartnerDetails[1]);
    let amountPartnerDetailsSet: {
      theatre: string;
      partner: string;
      amount: string;
    }[] = [];
    resultFilter[1].map((data) => {
      let amount =
        data["Minimum cost"] >
        parseInt(theatreFilter[1]["150"]) *
          parseInt(String(data["Cost Per GB"]))
          ? data["Minimum cost"]
          : parseInt(theatreFilter[1]["150"]) *
            parseInt(String(data["Cost Per GB"]));

      amountDetails = {
        ...amountDetails,
        theatre: data["Theatre"].trim(),
        deliverable: "true",
        partner: data["Partner ID"].trim(),
        amount: String(amount).trim(),
      };
      amountPartnerDetailsSet.push(amountDetails);
    });

    finalArrayList.push(amountPartnerDetailsSet[0]);
    let finalCsvArray: {
      distributor: string;
      deliverable: string;
      patner: string;
      amount: string;
    }[] = [];
    let finalCsvData = {
      distributor: "D4",
      deliverable: "false",
      patner: "''",
      amount: "''",
    };
    finalArrayList.push(finalCsvData);
    let totalAmount = 0;
    amountPartnerDetails.map((data) => {
      amountPartnerDetailsSet.map((inputData, index) => {
        if (index == 0) {
          totalAmount = parseInt(data.amount) + parseInt(inputData.amount);
          finalCsvData = {
            ...finalCsvData,
            deliverable: "true",
            patner: data.partner,
            amount: data.amount,
          };
        } else {
          if (
            totalAmount <
            parseInt(data.amount) + parseInt(inputData.amount)
          ) {
            finalCsvData = {
              ...finalCsvData,
              deliverable: "true",
              patner: data.partner,
              amount: data.amount,
            };
          }
        }
      });
    });
    csvParseService.writeCsvData("./output2.csv", finalArrayList);
  };
})();

export default appService;
