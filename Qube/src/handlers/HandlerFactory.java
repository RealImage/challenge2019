package handlers;

public class HandlerFactory
{
    public static Handler getHandlerObject(String type)
    {
        switch (type)
        {
            case "Theatre":
                return TheatreDetailHandler.getInstance();
            case "PartnerCapacity":
                return PartnerCapacityHandler.getInstance();
            case "PartnerDetails":
                return PartnerDetailsHandler.getInstance();
            case "InputDetails":
                return InputDetailHandler.getInstance();
            default:
                return null;
        }
    }
}
