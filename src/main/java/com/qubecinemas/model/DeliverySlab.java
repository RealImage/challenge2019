package com.qubecinemas.model;

/**
 * 
 * @author mohan
 *
 *DeliverySlab is hav the partner delivery slab, range and amount.
 */
public class DeliverySlab {

	private static int index = 1;

	public DeliverySlab(int minimumCost, int costPerGB, int startRange, int endRange, String partnerName) {
		this.id = index;
		this.minimumCost = minimumCost;
		this.costPerGB = costPerGB;
		this.startRange = startRange;
		this.endRange = endRange;
		this.partnerName = partnerName;
		index++;
	}

	private int id;

	private int minimumCost;

	private int costPerGB;

	private int startRange;

	private int endRange;

	private String partnerName;

	public int getId() {
		return id;
	}

	public void setId(int id) {
		this.id = id;
	}

	public int getMinimumCost() {
		return minimumCost;
	}

	public void setMinimumCost(int minimumCost) {
		this.minimumCost = minimumCost;
	}

	public int getCostPerGB() {
		return costPerGB;
	}

	public void setCostPerGB(int costPerGB) {
		this.costPerGB = costPerGB;
	}

	public int getStartRange() {
		return startRange;
	}

	public void setStartRange(int startRange) {
		this.startRange = startRange;
	}

	public int getEndRange() {
		return endRange;
	}

	public void setEndRange(int endRange) {
		this.endRange = endRange;
	}

	public String getPartnerName() {
		return partnerName;
	}

	public void setPartnerName(String partnerName) {
		this.partnerName = partnerName;
	}

	public boolean matchRange(int orderedGB) {
		return startRange < orderedGB && orderedGB <= endRange;
	}
}
