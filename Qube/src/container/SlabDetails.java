package container;

public class SlabDetails
{
    int minSlab;
    int maxSlab;
    int costPerGB;
    int minimumRate;

    public SlabDetails(int min, int max, int minimumRate, int costPerGB)
    {
        this.minSlab = min;
        this.maxSlab = max;
        this.costPerGB = costPerGB;
        this.minimumRate = minimumRate;
    }

    public int getMinSlab() {
        return minSlab;
    }

    public int getMaxSlab() {
        return maxSlab;
    }

    public int getMinimumRate() {
        return minimumRate;
    }

    public int getCostPerGB() {
        return costPerGB;
    }
}
