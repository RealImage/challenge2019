package com.realimage.csv.consumer;

import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

import com.opencsv.CSVWriter;
import com.opencsv.bean.CsvToBeanBuilder;
import com.realimage.bean.OutPut;

public class CsvUtility {

	@SuppressWarnings({ "unchecked", "rawtypes" })
	public static <T> List<T> consumeCSV(FileReader reader, Class<?> clazz) {
		List<T> list = new ArrayList<>();
		list = new CsvToBeanBuilder(reader).withType(clazz).build().parse();
		return list;
	}

	public static String getAbsolutPathofCSV(String filePath) {
		File file = new File(filePath);
		return file.getAbsolutePath();
	}

	public static boolean writeDataToCSV(List<OutPut> list, String filePath, String fileName) {
		File file = new File(filePath.concat("/").concat(fileName));
		try {
			CSVWriter csvWriter = new CSVWriter(new FileWriter(file));
			for (OutPut op : list) {
				String[] row = { op.getDeliveryId(), String.valueOf(op.isDeliveryPossible()),
						op.getPartnetId() != null ? op.getPartnetId() : "",
						op.getCostOfDelivery() > 0 ? String.valueOf(op.getCostOfDelivery()) : "" };
				csvWriter.writeNext(row);
			}
			csvWriter.close();

		} catch (IOException e) {
			e.printStackTrace();
			return false;
		}

		return true;
	}

}
