import { csvParseService } from "./CsvParseService";
import express from "express";

const app = express();

app.get("/", async (req, res) => {
  await csvParseService.parseCsvData("./input.csv");
  res.send("Hello");
});

app.listen(3000, () => {
  console.log("QUBE SERVER STARTED");
});
