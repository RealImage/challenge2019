package app;

import java.util.List;
import java.util.Map;

/**
 * DeliveryHandler
 */
public class DeliveryHandler {

    private Dataset datasource;

    private Map<String,List<Partner>> dataset;

    private List<Order> orders;

    public DeliveryHandler(Dataset datasource){
        this.datasource = datasource;
        dataset = datasource.getDataset();
        System.out.println(dataset);
        orders = datasource.getOrders();
        System.out.println(orders);
    }

    public List<Order> checkDeliveryOrders(){
        for (Order order : orders) {
            if(dataset.containsKey(order.getTheatreId())){
                assignPartner(order);
            }
        }
        return orders;
    }

    private void assignPartner(Order order) {

        List<Partner> partners = dataset.get(order.getTheatreId());
        int maxCost = Integer.MAX_VALUE;
        for (Partner partner : partners) {
            //System.out.println("isPossibleDelivery : "+partner.isPossibleDelivery(order));
            if(partner.isPossibleDelivery(order)){
                //System.out.println(" deliveryCharge : " + partner.deliveryCharge(order));
                if(partner.deliveryCharge(order) < maxCost){
                    maxCost = partner.deliveryCharge(order);
                    order.setPartner(partner);
                }
            }
        }
    }
    
}