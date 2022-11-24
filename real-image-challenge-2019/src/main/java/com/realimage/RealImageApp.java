package com.realimage;

import java.io.FileNotFoundException;
import java.io.FileReader;
import java.util.List;

import com.realimage.bean.Input;
import com.realimage.bean.Partner;
import com.realimage.csv.consumer.CsvUtility;
import com.realimage.service.RealImageAppService;
import com.realimage.service.impl.RealImageAppServiceImpl;

public class RealImageApp {

	RealImageAppService realImageAppService = new RealImageAppServiceImpl();

	// Files are saved to the resources folder
	static final String PARTNERS_FILE_PATH = CsvUtility.getAbsolutPathofCSV("src/main/resources/partners.csv");
	static final String INPUT_FILE_PATH = CsvUtility.getAbsolutPathofCSV("src/main/resources/input.csv");
	final String OUTPUT_1_PATH = CsvUtility.getAbsolutPathofCSV("src/main/resources/");
	final String OUTPUT_FILE_NAME = "output12.csv";

	List<Partner> partners;

	public static void main(String[] args) {
		List<Partner> partners;
		List<Input> inputs;
		try {
			partners = CsvUtility.consumeCSV(new FileReader(PARTNERS_FILE_PATH), Partner.class);
			inputs = CsvUtility.consumeCSV(new FileReader(INPUT_FILE_PATH), Input.class);
			RealImageApp realImageApp = new RealImageApp();
			realImageApp.init(inputs, partners);
		} catch (FileNotFoundException e) {
			e.printStackTrace();
		}

	}

	// File path and file name are passed to the service method to generate the CSV file
	public void init(List<Input> inputs, List<Partner> partners) {
		this.partners = realImageAppService.sortPartersByTheatreAndSlab(partners);
		if (this.partners != null && this.partners.size() > 0) {
			realImageAppService.findAndGenerateReportOfPartnersWithLeastDeliveryCost(inputs, this.partners,
					OUTPUT_1_PATH, OUTPUT_FILE_NAME);
		}

	}

}
