import fs from "fs";
import csv from "csv-parser";
import appService from "./AppService";
import ObjectsToCsv from "objects-to-csv";
import { Console } from "console";
export const csvParseService = new (class CsvParseService {
  public parseCsvData = async (fileName: String) => {
    const results: any[] = [];
    const resultData = await fs
      .createReadStream("./partners.csv")
      .pipe(csv({}))
      .on("data", (data) => results.push(data))
      .on("end", async () => {
        await appService.validateAndSolveProblemOne(results);
      });

    return results;
  };
  public writeCsvData = async (fileName: string, data: any[]) => {
    const csv = new ObjectsToCsv(data);
    await csv.toDisk(fileName);
  };
})();
