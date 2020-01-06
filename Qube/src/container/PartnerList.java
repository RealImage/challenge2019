package container;

import java.util.HashMap;
import java.util.Map;

public class PartnerList<T> implements ObjectList
{
    Map<String, T> partnerList = new HashMap();

    public Map<String, T> getList()
    {
        return partnerList;
    }
}
