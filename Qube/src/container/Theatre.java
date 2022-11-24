package container;

import java.util.HashMap;
import java.util.Map;

public class Theatre
{
    String theatreName;
    Map<String, Partner> partners;
    public Theatre(String theatreName)
    {
        this.theatreName = theatreName;
        this.partners = new HashMap<>();
    }

    public void addDetails()
    {

    }

    public Partner getPartner(String partnerName)
    {
        if(partners.containsKey(partnerName))
        {
            return partners.get(partnerName);
        }
        Partner p = new Partner(partnerName);
        partners.put(partnerName,p);
        return p;
    }

    public Map<String, Partner> getPartners()
    {
        return partners;
    }
}
