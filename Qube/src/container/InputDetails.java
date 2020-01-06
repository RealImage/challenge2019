package container;

public class InputDetails {

    String theatreName;
    String c1;
    int neededGB;
    boolean isPossible;
    public InputDetails(String c1, int neededGB, String theatreName)
    {
        this.theatreName = theatreName;
        this.c1 = c1;
        this.neededGB = neededGB;
    }

    public String getTheatreName() {
        return theatreName;
    }

    public int getNeededGB() {
        return neededGB;
    }

    public String getC1() {
        return c1;
    }

    public boolean isPossible()
    {
        return isPossible;
    }

    public void setPossible(boolean possible) {
        isPossible = possible;
    }
}
