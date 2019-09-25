package com.qubecinemas.model;

/**
 * 
 * @author mohan
 *
 *
 *         Delivery will requested delivery details.
 */
public class Delivery {

	public Delivery(String name, String theatre, int orderedGB) {
		this.name = name;
		this.theatre = theatre;
		this.orderedGB = orderedGB;
	}

	private String name;

	private String theatre;

	private int orderedGB;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getTheatre() {
		return theatre;
	}

	public void setTheatre(String theatre) {
		this.theatre = theatre;
	}

	public int getOrderedGB() {
		return orderedGB;
	}

	public void setOrderedGB(int requestedGB) {
		this.orderedGB = requestedGB;
	}

}
