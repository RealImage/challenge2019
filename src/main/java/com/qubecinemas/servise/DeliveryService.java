package com.qubecinemas.servise;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.qubecinemas.model.Delivery;
import com.qubecinemas.model.DeliveryDetail;
import com.qubecinemas.model.DeliveryManager;

/**
 * 
 * @author mohan
 * DeliveryService will help to deliver the all deliveries and will consider the partners capacity.
 */
public class DeliveryService {

	private TheatreService theatreService;

	public DeliveryService() {
		this.theatreService = new TheatreService();
	}

	/**
	 * 
	 * @param deliveries
	 * @return will find suitable partners to deliver the all deliveries and will consider the partners capacity.
	 */
	public Map<String, DeliveryDetail> getDeliveryDetails(List<Delivery> deliveries) {
		DeliveryManager deliveryManager = new DeliveryManager();
		deliveryManager.setDeliveries(deliveries);
		deliveryManager.setDeliveryDetails(findSuitablePartners(deliveries));
		while (deliveryManager.isHasCapacityConstraint()) {
			swapDeliveries(deliveryManager);
		}

		return deliveryManager.getDeliveryDetails();
	}

	private Map<String, DeliveryDetail> findSuitablePartners(List<Delivery> deliveries) {
		Map<String, DeliveryDetail> deliveryDetails = new HashMap<>();
		for (Delivery delivery : deliveries) {
			deliveryDetails.put(delivery.getName(), theatreService.getDeliveryPartner(delivery));
		}
		return deliveryDetails;
	}
	
	/**
	 * @param deliveryManager
	 * Will try to assign the deliveries to other partners who can deliver without capacity constrains. 
	 * If not will make delivery as cannot deliver.
 	 */
	private void swapDeliveries(DeliveryManager deliveryManager) {
		for (String partner : deliveryManager.getConstraintPartners()) {
			Map<String, DeliveryDetail> deliveries = deliveryManager.getDeliveryByPartner(partner);
			Map<String, DeliveryDetail> newDeliveries = new HashMap<>();
			for (DeliveryDetail detail : deliveries.values()) {
				DeliveryDetail newDelivery = detail;
				while (newDelivery == null || newDelivery.isCanDeliver()) {
					newDelivery = theatreService.getDeliveryPartner(
							deliveryManager.getDeliveries().get(detail.getDeliveryName()), newDelivery.getAmount());
					if (newDelivery.isCanDeliver() && deliveryManager.hasCapacity(newDelivery)) {
						break;
					}
				}
				newDeliveries.put(detail.getDeliveryName(), newDelivery);
			}
			getSuitableDeliveryPartner(deliveryManager, newDeliveries);
		}
	}

	/**
	 * 
	 * @param deliveryManager
	 * @param newDeliveries
	 * Will assign the delivery to next suitable partners and try to minimize the total cost.
	 */
	private void getSuitableDeliveryPartner(DeliveryManager deliveryManager,
			Map<String, DeliveryDetail> newDeliveries) {
		int currentTotal = Integer.MAX_VALUE;
		DeliveryDetail suitableDelivery = null;
		for (Map.Entry<String, DeliveryDetail> newDelivery : newDeliveries.entrySet()) {
			if (newDelivery.getValue().isCanDeliver()) {
				int exceptTotal = deliveryManager.getTotalCost()
						- deliveryManager.getDeliveryDetails().get(newDelivery.getKey()).getAmount();
				if (currentTotal > exceptTotal + newDelivery.getValue().getAmount()) {
					currentTotal = exceptTotal + newDelivery.getValue().getAmount();
					suitableDelivery = newDelivery.getValue();
				}
			}
		}
		deliveryManager.setDeliveryDetails(suitableDelivery);
	}

}
