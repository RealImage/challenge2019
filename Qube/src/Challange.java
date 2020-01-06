import container.*;
import handlers.HandlerFactory;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.atomic.AtomicReference;

public class Challange
{
    public static void main(String args[]) throws CloneNotSupportedException {
        ObjectList<Partner> partnerList = new PartnerList<>();
        ObjectList<InputDetails> inputList = new InputList();

        FileHandling.createObjectFromFile("partners.csv", false, partnerList, HandlerFactory.getHandlerObject("PartnerDetails"));
        FileHandling.createObjectFromFile("capacities.csv", false, partnerList, HandlerFactory.getHandlerObject("PartnerCapacity"));
        FileHandling.createObjectFromFile("input.csv", true, inputList, HandlerFactory.getHandlerObject("InputDetails"));

        GetResult getResult = new GetResult();
        ObjectList<InputDetails> possibleInputs = new InputList();
        ObjectList<InputDetails> imPossibleInputs = new InputList();
        getResult.getOutput1(inputList, partnerList, possibleInputs, imPossibleInputs);

        AtomicInteger totalAmount = new AtomicInteger(Integer.MAX_VALUE);
        List<OutputDetail> result = new ArrayList<>();


        getResult.getOutput2(((List<InputDetails>)possibleInputs.getList()), ((Map<String, Partner>)partnerList.getList()), totalAmount, result);

        StringBuilder s = new StringBuilder();
        result.stream().forEach(out ->
        {
            s.append(out).append(System.getProperty("line.separator"));
        });
        ((List<InputDetails>)imPossibleInputs.getList()).stream().forEach(imPossibleInput ->
        {
            s.append(imPossibleInput.getC1()).append(",").append(imPossibleInput.getTheatreName()).append(",").append("\"\"").append(",").append("\"\"").append(System.getProperty("line.separator"));
        });
        System.out.println(s);
        FileHandling.writeObjectToFile("out2.csv",s);
    }
}
