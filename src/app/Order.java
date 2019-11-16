package app;

/**
 * Order
 */
public class Order {

    private String orderId;

    private int size;

    private String theatreId;

    private Partner partner;

    public Order(String orderId,int size,String theatreId){
        this.orderId = orderId;
        this.size = size;
        this.theatreId = theatreId;
    }

    /**
     * @param orderId the orderId to set
     */
    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }

    /**
     * @param partner the partner to set
     */
    public void setPartner(Partner partner) {
        this.partner = partner;
    }

    /**
     * @param size the size to set
     */
    public void setSize(int size) {
        this.size = size;
    }

    /**
     * @param theatreId the theatreId to set
     */
    public void setTheatreId(String theatreId) {
        this.theatreId = theatreId;
    }

    /**
     * @return the orderId
     */
    public String getOrderId() {
        return orderId;
    }

    /**
     * @return the partner
     */
    public Partner getPartner() {
        return partner;
    }

    /**
     * @return the size
     */
    public int getSize() {
        return size;
    }

    /**
     * @return the theatreId
     */
    public String getTheatreId() {
        return theatreId;
    }

    public boolean isDeliveryPossible(){
        return partner == null ? false : true;
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("[").append(orderId).append(",")
        .append(this.size).append(",")
        .append(this.theatreId).append(",")
        .append(this.isDeliveryPossible());
        if(this.isDeliveryPossible()){
            sb.append(",").append(this.partner.deliveryCharge(this));
            sb.append(",").append(this.partner.getId());
        }
        sb.append("]");
        return sb.toString();

    }

}