package com.qubecinemas.model;

/**
 * 
 * @author mohan
 *
 *         Capacity is data model to hold Partner capacity.
 */
public class Capacity {

	public Capacity(String partner, int capacity) {
		this.partner = partner;
		this.capacity = capacity;
	}

	private String partner;

	private int capacity;

	public String getPartner() {
		return partner;
	}

	public void setPartner(String partner) {
		this.partner = partner;
	}

	public int getCapacity() {
		return capacity;
	}

	public void setCapacity(int capacity) {
		this.capacity = capacity;
	}

}
