package container;

import java.util.HashMap;
import java.util.Map;

public class TheatreList<T> implements ObjectList
{
    Map<String, T> theatreList = new HashMap();

    public Map<String, T> getList()
    {
        return theatreList;
    }

}
