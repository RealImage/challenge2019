package handlers;

import container.InputDetails;
import container.ObjectList;

import java.util.List;

public class InputDetailHandler<T> implements Handler
{

    static InputDetailHandler id = null;

    public static InputDetailHandler getInstance()
    {
        if(id == null)
        {
            id = new InputDetailHandler();
        }
        return id;
    }

    @Override
    public void handleDetails(String[] details, ObjectList list)
    {
        ((List)list.getList()).add(new InputDetails(details[0].trim(), Integer.parseInt(details[1].trim()), details[2].trim()));
    }
}
