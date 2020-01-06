package container;

public class OutputDetail implements Cloneable
{
    String partnerName;
    int totalRate;
    String c1;
    boolean isPossible;
    int totalGB;

    public String getC1() {
        return c1;
    }

    public String getPartnerName() {
        return partnerName;
    }

    public int getTotalRate() {
        return totalRate;
    }

    public int getTotalGB() {
        return totalGB;
    }

    public void setTotalGB(int totalGB) {
        this.totalGB = totalGB;
    }

    public void setC1(String c1) {
        this.c1 = c1;
    }

    public void setPartnerName(String partnerName) {
        this.partnerName = partnerName;
    }

    public void setPossible(boolean possible) {
        isPossible = possible;
    }

    public void setTotalRate(int totalRate) {
        this.totalRate = totalRate;
    }

    @Override
    public Object clone() throws CloneNotSupportedException {
        return super.clone();
    }

    public void setDetails(InputDetails inputDetails, String partnerName, int totalRate, int neededGB)
    {
        this.setC1(inputDetails.getC1());
        this.setPossible(true);
        this.setTotalRate(totalRate);
        this.setPartnerName(partnerName);
        this.setTotalGB(neededGB);
    }

    @Override
    public String toString() {
        return this.getC1()+","+this.isPossible+","+this.getPartnerName()+","+this.getTotalRate();
    }
}
