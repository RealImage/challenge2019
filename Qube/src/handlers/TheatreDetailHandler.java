package handlers;

import container.ObjectList;
import container.Partner;
import container.Theatre;
import container.TheatreList;

import java.util.Map;

public class TheatreDetailHandler implements Handler
{
    static TheatreDetailHandler th = null;
    public static TheatreDetailHandler getInstance()
    {
        if(th == null)
        {
            th = new TheatreDetailHandler();
        }
        return th;
    }

    @Override
    public void handleDetails(String[] theatreDetails, ObjectList obj)
    {

    }
}
