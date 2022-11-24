package com.realimage.service;

import java.util.List;

import com.realimage.bean.Input;
import com.realimage.bean.Partner;

public interface RealImageAppService {

	List<Partner> sortPartersByTheatreAndSlab(List<Partner> partners);

	void findAndGenerateReportOfPartnersWithLeastDeliveryCost(List<Input> inputs, List<Partner> partners,
			String outputPath, String fileName);

}
