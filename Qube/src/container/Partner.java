package container;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Partner
{
    String partnerName;
    int capacity;
    Map<String, List<SlabDetails>> theatreSlabDetails;
    public Partner(String partnerName)
    {
        this.partnerName = partnerName;
        this.theatreSlabDetails = new HashMap<>();
    }

    public void addDetails(String theatreName,String slabRange, String minimumRate, String costPerGB)
    {
        List slabDetails = theatreSlabDetails.computeIfAbsent(theatreName, k-> new ArrayList<>());
        String[] range = slabRange.split("-");
        slabDetails.add(new SlabDetails(Integer.parseInt(range[0].trim()), Integer.parseInt(range[1].trim()), Integer.parseInt(minimumRate), Integer.parseInt(costPerGB)));
    }

    public String getPartnerName() {
        return partnerName;
    }

    public Map<String, List<SlabDetails>> getTheatreSlabDetails() {
        return theatreSlabDetails;
    }

    public int getCapacity() {
        return capacity;
    }

    public void setCapacity(int capacity) {
        this.capacity = capacity;
    }
}
