package app;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Dataset
 */
public class Dataset {

    private static Map<String, List<Partner>> dataset;

    private static List<Order> orders;

    private final String DATASET_FILENAME = "partners.csv";

    private final String INPUT_FILENAME = "input.csv";

    public Map<String, List<Partner>> getDataset() {
        if(dataset != null && !dataset.isEmpty()) return dataset;
        dataset = new HashMap<String,List<Partner>>();
        buildDataset();
        return dataset;
    }

    public List<Order> getOrders(){
        if(orders != null && !orders.isEmpty()) return orders;
        orders = new ArrayList<Order>();
        buildOrders();
        return orders;
    }

    private List<Order> buildOrders(){
        new Parser() {

            @Override
            protected void mapper(String[] values) {
                Order order = createOrder(values);
                orders.add(order);
            }
        }.parseCSV(INPUT_FILENAME, ",",false);;

        return orders;
    }

    private Map<String, List<Partner>> buildDataset() {
        new Parser() {

            @Override
            protected void mapper(String[] values) {
                Partner partner = createPartner(values);
                if (dataset.containsKey(values[0].trim())) {
                    List<Partner> partners = dataset.get(values[0].trim());
                    partners.add(partner);
                } else {
                    List<Partner> partners = new ArrayList<Partner>();
                    partners.add(partner);
                    dataset.put(values[0].trim(), partners);
                }
            }
        }.parseCSV(DATASET_FILENAME, ",",true);;

        return dataset;
    }

    private Partner createPartner(String[] data) {
        String[] size = data[1].split("-");
        int minCost = Integer.parseInt(data[2].trim());
        int mn = Integer.parseInt(size[0].trim());
        int mx = Integer.parseInt(size[1].trim());
        int cost = Integer.parseInt(data[3].trim());
        return new Partner(data[4].trim(), mn, mx, cost,minCost);
    }

    private Order createOrder(String[] data) {
        String id = data[0].trim();
        int size = Integer.parseInt(data[1].trim());
        String theatreId = data[2].trim();
        return new Order(id, size, theatreId);
    }

}