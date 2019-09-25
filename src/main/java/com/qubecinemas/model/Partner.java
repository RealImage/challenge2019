package com.qubecinemas.model;

import java.util.Set;

/**
 * 
 * @author mohan
 * Partner will contain all delivery slab and cost details.
 */
public class Partner {

	
	public Partner(String name, Set<DeliverySlab> deliverySlabs) {
		this.name = name;
		this.deliverySlabs = deliverySlabs;
	}

	public Partner() {
	}

	private String name;

	private Set<DeliverySlab> deliverySlabs;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public Set<DeliverySlab> getDeliverySlabs() {
		return deliverySlabs;
	}

	public void setDeliverySlabs(Set<DeliverySlab> deliverySlabs) {
		this.deliverySlabs = deliverySlabs;
	}
}
