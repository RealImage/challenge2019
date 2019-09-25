package com.qubecinemas.model;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Set;

/**
 * 
 * @author mohan
 * Theatre will have partners who can delivery the order.
 */
public class Theatre {

	public Theatre(String name, Set<Partner> partners) {
		this.name = name;
		this.partners = partners;
	}

	public Theatre() {
	}

	private String name;

	private Set<Partner> partners;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public Set<Partner> getPartners() {
		if (partners == null)
			return Collections.emptySet();
		return partners;
	}

	public void setPartners(Set<Partner> partners) {
		this.partners = partners;
	}

	public List<DeliverySlab> getPartnersDeliverySlabs() {
		List<DeliverySlab> deliverySlabs = new ArrayList<>();
		for (Partner partner : getPartners()) {
			deliverySlabs.addAll(partner.getDeliverySlabs());
		}
		return deliverySlabs;
	}
}
