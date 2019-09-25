package com.qubecinemas;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

import com.qubecinemas.data.DataSet;
import com.qubecinemas.model.Delivery;
import com.qubecinemas.model.DeliveryDetail;
import com.qubecinemas.servise.DeliveryService;

import junit.framework.TestCase;

public class TestProblem_2 extends TestCase {

	DeliveryService service = null;

	public void setUp() {
		DataSet.loadData();
		service = new DeliveryService();
	}

	public void testCapacityDelivery() {
		List<Delivery> deliveries = new ArrayList<>();
		Delivery D1 = new Delivery("D1", "T1", 150);
		Delivery D2 = new Delivery("D2", "T2", 325);
		Delivery D3 = new Delivery("D3", "T1", 510);
		Delivery D4 = new Delivery("D4", "T2", 700);
		deliveries.add(D1);
		deliveries.add(D2);
		deliveries.add(D3);
		deliveries.add(D4);

		Map<String, DeliveryDetail> deliveryDetails = service.getDeliveryDetails(deliveries);
	
		assertEquals(true, deliveryDetails.get("D1").isCanDeliver());
		assertEquals("P1", deliveryDetails.get("D1").getPartnerName());
		assertEquals(2000, deliveryDetails.get("D1").getAmount());
		
		assertEquals(true, deliveryDetails.get("D2").isCanDeliver());
		assertEquals("P2", deliveryDetails.get("D2").getPartnerName());
		assertEquals(3500, deliveryDetails.get("D2").getAmount());
		
		
		assertEquals(true, deliveryDetails.get("D3").isCanDeliver());
		assertEquals("P3", deliveryDetails.get("D3").getPartnerName());
		assertEquals(15300, deliveryDetails.get("D3").getAmount());
		
		assertEquals(false, deliveryDetails.get("D4").isCanDeliver());
		assertEquals("", deliveryDetails.get("D4").getPartnerName());
		assertEquals(0, deliveryDetails.get("D4").getAmount());

	}
}
