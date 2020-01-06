package handlers;

import container.ObjectList;
import container.Partner;

import java.util.Map;

public class PartnerCapacityHandler<T> implements Handler
{
    static PartnerCapacityHandler pc = null;
    public static PartnerCapacityHandler getInstance()
    {
        if(pc == null)
        {
           pc = new PartnerCapacityHandler();
        }
        return pc;
    }
    @Override
    public void handleDetails(String[] details, ObjectList partnerList)
    {
        Partner partner = (Partner) ((Map<String, Partner>)partnerList.getList()).get(details[0].trim());
        partner.setCapacity(Integer.parseInt(details[1].trim()));

    }
}
