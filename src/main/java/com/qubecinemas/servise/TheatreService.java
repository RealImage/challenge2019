package com.qubecinemas.servise;

import java.util.Set;
import java.util.stream.Collectors;

import com.qubecinemas.data.DataSet;
import com.qubecinemas.model.Delivery;
import com.qubecinemas.model.DeliveryDetail;
import com.qubecinemas.model.DeliverySlab;
import com.qubecinemas.model.Partner;
import com.qubecinemas.model.Theatre;

/**
 * 
 * @author mohan
 * TheatreService will provide the function to find the suitable partner who can deliver the order.
 */
public class TheatreService {

	/**
	 * 
	 * @param theatreName
	 * @param orderedGB
	 * @return the eligible delivery slab which can deliver based on theatre name and ordered GB.
	 */
	private Set<DeliverySlab> getPartnerSlabs(String theatreName, int orderedGB) {
		Theatre theatre = DataSet.theatreList.get(theatreName);

		return theatre.getPartnersDeliverySlabs().stream().filter(slab -> slab.matchRange(orderedGB))
				.collect(Collectors.toSet());

	}

	/**
	 * 
	 * @param orderedGB
	 * @param eligiblePartners
	 * @param deliveryDetails
	 * @param minimumLimit
	 * 
	 * Will find the suitable {@link Partner} based on the slab and minimumlimit.
	 */
	private void findSuitablePartner(int orderedGB, Set<DeliverySlab> eligiblePartners,
			DeliveryDetail deliveryDetails, int minimumLimit) {
		int currentMIn = Integer.MAX_VALUE;
		for (DeliverySlab slab : eligiblePartners) {
			int cost = orderedGB * slab.getCostPerGB();
			if (cost < slab.getMinimumCost())
				cost = slab.getMinimumCost();
			if (currentMIn > cost && cost > minimumLimit) {
				deliveryDetails.setCanDeliver(true);
				deliveryDetails.setAmount(cost);
				deliveryDetails.setPartnerName(slab.getPartnerName());
				currentMIn = cost;
			}
		}
	}

	public DeliveryDetail getDeliveryPartner(Delivery delivery) {
		return getDeliveryPartner(delivery, 0);
	}
	
	/**
	 * 
	 * @param delivery
	 * @param minimumLimit
	 * @return Get the suitable order which can be deliver and if not found will mark delivery as cannot deliver.
	 */
	public DeliveryDetail getDeliveryPartner(Delivery delivery, int minimumLimit) {
		DeliveryDetail deliveryDetails = new DeliveryDetail();
		deliveryDetails.setDeliveryName(delivery.getName());
		deliveryDetails.setTheatreName(delivery.getTheatre());
		deliveryDetails.setOrderedGB(delivery.getOrderedGB());
		Set<DeliverySlab> eligiblePartners = getPartnerSlabs(delivery.getTheatre(), delivery.getOrderedGB());
		if (eligiblePartners.isEmpty()) {
			deliveryDetails.setCanDeliver(false);
		} else {
			findSuitablePartner(delivery.getOrderedGB(), eligiblePartners, deliveryDetails, minimumLimit);
		}
		return deliveryDetails;
	}
}
