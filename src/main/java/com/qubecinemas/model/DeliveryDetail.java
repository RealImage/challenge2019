package com.qubecinemas.model;

/**
 * 
 * @author mohan
 *
 *         DeliveryDetail is the data model to hold delivery partner details for
 *         each delivery
 */
public class DeliveryDetail {

	private String deliveryName;
	private boolean canDeliver;
	private String theatreName;
	private int orderedGB;
	private String partnerName;
	private int amount;

	public String getDeliveryName() {
		return deliveryName;
	}

	public void setDeliveryName(String deliveryName) {
		this.deliveryName = deliveryName;
	}

	public boolean isCanDeliver() {
		return canDeliver;
	}

	public void setCanDeliver(boolean canDeliver) {
		this.canDeliver = canDeliver;
	}

	public String getPartnerName() {
		if (partnerName == null)
			return "";
		return partnerName;
	}

	public void setPartnerName(String partnerName) {
		this.partnerName = partnerName;
	}

	public int getAmount() {
		return amount;
	}

	public void setAmount(int amount) {
		this.amount = amount;
	}

	public String getTheatreName() {
		return theatreName;
	}

	public void setTheatreName(String theatreName) {
		this.theatreName = theatreName;
	}

	public int getOrderedGB() {
		return orderedGB;
	}

	public void setOrderedGB(int orderedGB) {
		this.orderedGB = orderedGB;
	}

	@Override
	public String toString() {
		return "DeliveryDetails [deliveryName=" + deliveryName + ", canDeliver=" + canDeliver + ", partnerName="
				+ partnerName + ", amount=" + amount + "]";
	}

}
