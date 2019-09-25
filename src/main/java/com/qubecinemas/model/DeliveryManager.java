package com.qubecinemas.model;

import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

import com.qubecinemas.data.DataSet;

/**
 * 
 * @author mohan
 *
 *         DeliveryManager will manage the List of deliveries try to make all
 *         the order possible and try to reduce the total amount.
 * 
 */
public class DeliveryManager {

	private Map<String, DeliveryDetail> deliveryDetails;

	private Map<String, Delivery> deliveries;

	private boolean hasCapacityConstraint;

	private Set<String> constraintPartners;

	public Map<String, DeliveryDetail> getDeliveryDetails() {
		if (deliveryDetails == null)
			return Collections.emptyMap();
		return deliveryDetails;
	}

	public void setDeliveryDetails(Map<String, DeliveryDetail> deliveryDetails) {
		this.deliveryDetails = deliveryDetails;
	}

	/**
	 * 
	 * @param deliveryDetail
	 * 
	 * Add the new delivery to the delivery list.
	 */
	public void setDeliveryDetails(DeliveryDetail deliveryDetail) {
		if (this.deliveryDetails == null || this.deliveryDetails.isEmpty())
			this.deliveryDetails = new HashMap<>();
		this.deliveryDetails.put(deliveryDetail.getDeliveryName(), deliveryDetail);
	}

	/**
	 * 
	 * @return
	 * It will check all the delivery against partners and check they have any capacity constrains.
	 */
	
	public boolean isHasCapacityConstraint() {
		if (deliveryDetails.isEmpty()) {
			return true;
		}
		checkParnersCapacity();
		return hasCapacityConstraint;
	}

	public Set<String> getConstraintPartners() {
		if (constraintPartners == null)
			return Collections.emptySet();
		return constraintPartners;
	}

	/**
	 * 
	 * @return
	 * 
	 * Will calculate the total cost all delivery.
	 */
	public int getTotalCost() {
		return deliveryDetails.values().stream().collect(Collectors.summingInt(DeliveryDetail::getAmount));
	}

	public Map<String, Delivery> getDeliveries() {
		if (deliveries == null)
			return Collections.emptyMap();
		return deliveries;
	}

	public void setDeliveries(List<Delivery> deliveries) {
		Map<String, Delivery> deli = new HashMap<>();
		deli.putAll(getDeliveries());
		for (Delivery delivery : deliveries) {
			deli.put(delivery.getName(), delivery);
		}
		this.deliveries = deli;
	}

	/**
	 * It will check all the delivery against partners and check they have any capacity constrains.
	 */
	private void checkParnersCapacity() {
		hasCapacityConstraint = false;
		constraintPartners = new HashSet<>();
		Map<String, Integer> currentCapacity = deliveryDetails.values().stream()
				.filter(delivery -> delivery.isCanDeliver()).collect(Collectors.groupingBy(
						DeliveryDetail::getPartnerName, Collectors.summingInt(DeliveryDetail::getOrderedGB)));
		for (Map.Entry<String, Integer> partnerCap : currentCapacity.entrySet()) {
			if (DataSet.partnerCapacity.get(partnerCap.getKey()).getCapacity() < partnerCap.getValue()) {
				hasCapacityConstraint = true;
				constraintPartners.add(partnerCap.getKey());
			}
		}
	}

	/**
	 * 
	 * @param partner
	 * @return
	 * 
	 * returns deliverydetails by the partner name.
	 */
	public Map<String, DeliveryDetail> getDeliveryByPartner(String partner) {
		return deliveryDetails.values().stream().filter(delivery -> partner.equals(delivery.getPartnerName()))
				.collect(Collectors.toMap(DeliveryDetail::getDeliveryName, delivery -> delivery));
	}

	/**
	 * @param newDelivery
	 * @return
	 * 
	 * check if partner can deliver with this new order.
	 */
	public boolean hasCapacity(DeliveryDetail newDelivery) {
		int oldCapacity = getDeliveryByPartner(newDelivery.getPartnerName()).values().stream()
				.collect(Collectors.summingInt(DeliveryDetail::getOrderedGB));
		return DataSet.partnerCapacity.get(newDelivery.getPartnerName()).getCapacity() >= oldCapacity
				+ newDelivery.getOrderedGB();
	}
}
