package com.qubecinemas.data;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

import com.qubecinemas.model.Capacity;
import com.qubecinemas.model.DeliverySlab;
import com.qubecinemas.model.Partner;
import com.qubecinemas.model.Theatre;

/**
 * 
 * @author mohan
 *
 *         DataSet is act like Data source for problems. All these data is
 *         loaded from CSV files.
 */
public class DataSet {

	private DataSet() {
	}

	public static Map<String, Theatre> theatreList = new HashMap<String, Theatre>();

	public static Map<String, Capacity> partnerCapacity = new HashMap<String, Capacity>();

	public static void loadData() {
		loadT1();
		loadT2();
		loadCapacity();
	}

	private static void loadCapacity() {
		Capacity P1_C = new Capacity("P1", 350);
		Capacity P2_C = new Capacity("P2", 500);
		Capacity P3_C = new Capacity("P3", 1500);
		partnerCapacity.put(P1_C.getPartner(), P1_C);
		partnerCapacity.put(P2_C.getPartner(), P2_C);
		partnerCapacity.put(P3_C.getPartner(), P3_C);
	}

	/**
	 * Loading Theatre 1 data
	 */
	private static void loadT1() {
		Theatre T1 = new Theatre();

		Set<Partner> T1_PS = new HashSet<Partner>();

		Partner T1_P1 = new Partner();

		Set<DeliverySlab> T1_P1_DS = new HashSet<DeliverySlab>();
		T1_P1_DS.add(new DeliverySlab(1500, 20, 0, 100, "P1"));
		T1_P1_DS.add(new DeliverySlab(2000, 13, 100, 200, "P1"));
		T1_P1_DS.add(new DeliverySlab(2500, 12, 200, 300, "P1"));
		T1_P1_DS.add(new DeliverySlab(3000, 10, 300, 400, "P1"));

		T1_P1.setName("P1");
		T1_P1.setDeliverySlabs(T1_P1_DS);

		Partner T1_P2 = new Partner();

		Set<DeliverySlab> T1_P2_DS = new HashSet<DeliverySlab>();
		T1_P2_DS.add(new DeliverySlab(1000, 20, 0, 200, "P2"));
		T1_P2_DS.add(new DeliverySlab(2500, 15, 200, 400, "P2"));

		T1_P2.setName("P2");
		T1_P2.setDeliverySlabs(T1_P2_DS);

		Partner T1_P3 = new Partner();

		Set<DeliverySlab> T1_P3_DS = new HashSet<DeliverySlab>();
		T1_P3_DS.add(new DeliverySlab(800, 25, 100, 200, "P3"));
		T1_P3_DS.add(new DeliverySlab(1200, 30, 200, 600, "P3"));

		T1_P3.setName("P3");
		T1_P3.setDeliverySlabs(T1_P3_DS);

		T1_PS.add(T1_P1);
		T1_PS.add(T1_P2);
		T1_PS.add(T1_P3);

		T1.setName("T1");
		T1.setPartners(T1_PS);

		theatreList.put(T1.getName(), T1);
	}

	/**
	 * Loading Theatre T2 data
	 */
	private static void loadT2() {
		Theatre T2 = new Theatre();

		Set<Partner> T2_PS = new HashSet<Partner>();

		Partner T2_P1 = new Partner();

		Set<DeliverySlab> T2_P1_DS = new HashSet<DeliverySlab>();
		T2_P1_DS.add(new DeliverySlab(1500, 20, 0, 100, "P1"));
		T2_P1_DS.add(new DeliverySlab(2000, 15, 100, 200, "P1"));
		T2_P1_DS.add(new DeliverySlab(2500, 12, 200, 300, "P1"));
		T2_P1_DS.add(new DeliverySlab(3000, 10, 300, 400, "P1"));

		T2_P1.setName("P1");
		T2_P1.setDeliverySlabs(T2_P1_DS);

		Partner T2_P2 = new Partner();

		Set<DeliverySlab> T2_P2_DS = new HashSet<DeliverySlab>();
		T2_P2_DS.add(new DeliverySlab(2500, 20, 0, 200, "P2"));
		T2_P2_DS.add(new DeliverySlab(3500, 10, 200, 400, "P2"));

		T2_P2.setName("P2");
		T2_P2.setDeliverySlabs(T2_P2_DS);

		Partner T2_P3 = new Partner();

		Set<DeliverySlab> T2_P3_DS = new HashSet<DeliverySlab>();
		T2_P3_DS.add(new DeliverySlab(900, 15, 100, 200, "P3"));
		T2_P3_DS.add(new DeliverySlab(1000, 12, 200, 400, "P3"));

		T2_P3.setName("P3");
		T2_P3.setDeliverySlabs(T2_P3_DS);

		T2_PS.add(T2_P1);
		T2_PS.add(T2_P2);
		T2_PS.add(T2_P3);

		T2.setName("T2");
		T2.setPartners(T2_PS);

		theatreList.put(T2.getName(), T2);
	}
}
