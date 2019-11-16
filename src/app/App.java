package app;

import java.util.List;

public class App {
    public static void main(String[] args) throws Exception {
        DeliveryHandler deliveryHandler = new DeliveryHandler(new Dataset());
        List<Order> orders = deliveryHandler.checkDeliveryOrders();
        System.out.println(orders);
    }
}