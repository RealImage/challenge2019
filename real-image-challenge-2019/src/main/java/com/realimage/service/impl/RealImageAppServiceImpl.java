package com.realimage.service.impl;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.stream.Collectors;

import org.apache.commons.lang3.Range;

import com.realimage.bean.Input;
import com.realimage.bean.OutPut;
import com.realimage.bean.Partner;
import com.realimage.csv.consumer.CsvUtility;
import com.realimage.service.RealImageAppService;

public class RealImageAppServiceImpl implements RealImageAppService {

	@Override
	public List<Partner> sortPartersByTheatreAndSlab(List<Partner> partners) {
		Comparator<Partner> comparator = Comparator.comparing(Partner::getTheatre).thenComparing(Partner::getSizeSlab);
		return partners.stream().sorted(comparator).collect(Collectors.toList());

	}

	@Override
	public void findAndGenerateReportOfPartnersWithLeastDeliveryCost(List<Input> inputs, List<Partner> partners,
			String outputPath, String fileName) {
		List<OutPut> tempOutputs = null;
		List<OutPut> finalList = new ArrayList<>();
		OutPut output;

		for (Input input : inputs) {
			tempOutputs = new ArrayList<>();
			List<Partner> p1 = getPartnersByTheatreAndSlab(input.getTheatreId(), input.getSizeOfDelivery(), partners);
			if (p1 != null && p1.size() > 0) {
				for (Partner partner : p1) {
					if (partner.getTheatre().trim().equals(input.getTheatreId())) {
						output = new OutPut();
						String[] contentSizeRange = partner.getSizeSlab().split("-");
						Range<Integer> range = Range.between(Integer.parseInt(contentSizeRange[0].trim()),
								Integer.parseInt(contentSizeRange[1].trim()));
						if (range.contains(input.getSizeOfDelivery())) {
							int costOfDelivery = input.getSizeOfDelivery() * partner.getCostPerGB();
							costOfDelivery = costOfDelivery >= partner.getMinimumSlab() ? costOfDelivery
									: partner.getMinimumSlab();
							output.setCostOfDelivery(costOfDelivery);
							output.setDeliveryId(input.getDeliveryId());
							output.setDeliveryPossible(true);
							output.setPartnetId(partner.getPartnerId());
							tempOutputs.add(output);

						} else {
							output.setCostOfDelivery(0);
							output.setDeliveryId(input.getDeliveryId());
							output.setDeliveryPossible(false);
							output.setPartnetId(null);
							tempOutputs.add(output);

						}
					}

				}

			} else {
				output = new OutPut();
				output.setCostOfDelivery(0);
				output.setDeliveryId(input.getDeliveryId());
				output.setDeliveryPossible(false);
				output.setPartnetId(null);
				tempOutputs.add(output);

			}
			finalList.addAll(extractTheLeastOne(tempOutputs));

		}
		if (finalList != null && finalList.size() > 0) {
			CsvUtility.writeDataToCSV(finalList, outputPath, fileName);
		}

	}

	private List<Partner> getPartnersByTheatreAndSlab(String theatreId, int slab, List<Partner> partners) {
		List<Partner> result = new ArrayList<>();
		partners.forEach(p -> {
			String[] splittedSlab = p.getSizeSlab().split("-");
			int lowerSlab = Integer.parseInt(splittedSlab[0].trim());
			int upperSlab = Integer.parseInt(splittedSlab[1].trim());
			if (p.getTheatre().trim().equals(theatreId) && (lowerSlab <= slab && upperSlab >= slab)) {
				result.add(p);
			}
		});
		return result;
	}

	private List<OutPut> extractTheLeastOne(List<OutPut> tempOutputs) {
		List<OutPut> outputList = null;
		if (tempOutputs != null && tempOutputs.size() > 0) {
			Comparator<OutPut> comparator = Comparator.comparing(OutPut::getCostOfDelivery);
			List<OutPut> op = tempOutputs.stream().sorted(comparator).collect(Collectors.toList());
			outputList = new ArrayList<>();
			outputList.add(op.get(0));
			outputList.forEach(System.out::println);
		}
		return outputList;

	}
	

}
