package app;

/**
 * Partner
 */
public class Partner {

    private String id;

    private int min,max,cost,minCost;

    Partner(String id,int mn,int mx,int cost,int mCost){
        this.id = id;
        min = mn;
        max = mx;
        this.cost = cost;
        this.minCost = mCost;
    }

    /**
     * @param minCost the minCost to set
     */
    public void setMinCost(int minCost) {
        this.minCost = minCost;
    }

    /**
     * @return the minCost
     */
    public int getMinCost() {
        return minCost;
    }

    /**
     * @param cost the cost to set
     */
    public void setCost(int cost) {
        this.cost = cost;
    }

    /**
     * @return the cost
     */
    public int getCost() {
        return cost;
    }

    /**
     * @param id the id to set
     */
    public void setId(String id) {
        this.id = id;
    }

    /**
     * @return the id
     */
    public String getId() {
        return id;
    }

    /**
     * @param max the max to set
     */
    public void setMax(int max) {
        this.max = max;
    }

    /**
     * @param min the min to set
     */
    public void setMin(int min) {
        this.min = min;
    }

    /**
     * @return the max
     */
    public int getMax() {
        return max;
    }

    /**
     * @return the min
     */
    public int getMin() {
        return min;
    }

    public boolean isPossibleDelivery(Order order){
        return (order.getSize() > min && order.getSize() <= max);
    }

    public int deliveryCharge(Order order){
        return (order.getSize() * cost);
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("[").append(this.id).append(",")
        .append(this.min).append(",").append(this.max)
        .append(",").append(this.cost).append(",")
        .append(this.minCost).append("]");
        return sb.toString();
    }



}