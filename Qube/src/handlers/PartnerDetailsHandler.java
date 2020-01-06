package handlers;

import container.*;

import java.util.Map;

public class PartnerDetailsHandler<T> implements Handler
{

    static PartnerDetailsHandler pd = null;

    public static PartnerDetailsHandler getInstance()
    {
        if(pd == null)
        {
            pd = new PartnerDetailsHandler();
        }
        return pd;
    }
    @Override
    public void handleDetails(String[] partnerDetails, ObjectList pl)
    {
        String theatreName = partnerDetails[0].trim();
        String partnerName = partnerDetails[4].trim();
        Map<String, Partner> partnerList =(Map<String, Partner>) pl.getList();
        Partner partner;
        if(partnerList.containsKey(partnerName))
        {
            partner = partnerList.get(partnerName);
        }
        else
        {
            partner = new Partner(partnerName);
            partnerList.put(partnerName, partner);
        }
        partner.addDetails(theatreName,partnerDetails[1].trim(), partnerDetails[2].trim(), partnerDetails[3].trim());

    }
}
