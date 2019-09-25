package com.qubecinemas;

import com.qubecinemas.data.DataSet;
import com.qubecinemas.model.Delivery;
import com.qubecinemas.model.DeliveryDetail;
import com.qubecinemas.servise.TheatreService;

import junit.framework.TestCase;

public class TestProblem_1 extends TestCase {

	TheatreService service = null;

	public void setUp() {
		DataSet.loadData();
		service = new TheatreService();
	}

	public void testDelivery1() {
		Delivery delivery = new Delivery("D1", "T1", 150);

		DeliveryDetail details = service.getDeliveryPartner(delivery);

		assertEquals(true, details.isCanDeliver());
		assertEquals("P1", details.getPartnerName());
		assertEquals(2000, details.getAmount());
	}

	public void testDelivery2() {
		Delivery delivery = new Delivery("D2", "T2", 325);

		DeliveryDetail details = service.getDeliveryPartner(delivery);

		assertEquals(true, details.isCanDeliver());
		assertEquals("P1", details.getPartnerName());
		assertEquals(3250, details.getAmount());
	}

	public void testDelivery3() {
		Delivery delivery = new Delivery("D3", "T1", 510);

		DeliveryDetail details = service.getDeliveryPartner(delivery);

		assertEquals(true, details.isCanDeliver());
		assertEquals("P3", details.getPartnerName());
		assertEquals(15300, details.getAmount());
	}

	public void testDelivery4() {
		Delivery delivery = new Delivery("D4", "T2", 700);

		DeliveryDetail details = service.getDeliveryPartner(delivery);

		assertEquals(false, details.isCanDeliver());
		assertEquals("", details.getPartnerName());
		assertEquals(0, details.getAmount());
	}
}
