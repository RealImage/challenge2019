import container.*;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.atomic.AtomicReference;

public class GetResult
{

    public void getOutput1(ObjectList<InputDetails> inputList, ObjectList<Partner> partnerList, ObjectList<InputDetails> possibleInputs, ObjectList<InputDetails> imPossibleInputs)
    {
        StringBuilder sb = new StringBuilder();
        ((List<InputDetails>)inputList.getList()).stream().forEach(k->
        {
            sb.append(k.getC1()).append(",");
            AtomicInteger min = new AtomicInteger(Integer.MAX_VALUE);
            AtomicReference<String> partnerName = new AtomicReference<>("");
            ((Map<String, Partner>)partnerList.getList()).entrySet().stream().forEach(partnerEntry ->
            {
                List<SlabDetails> list = partnerEntry.getValue().getTheatreSlabDetails().get(k.getTheatreName());
                list.stream().forEach(slabDetails ->
                {
                    if(slabDetails.getMinSlab() <= k.getNeededGB() && slabDetails.getMaxSlab() >= k.getNeededGB())
                    {
                        int maxRate = slabDetails.getCostPerGB() * k.getNeededGB();
                        if(maxRate < slabDetails.getMinimumRate())
                        {
                            maxRate = slabDetails.getMinimumRate();
                        }
                        if(maxRate < min.get())
                        {
                            min.set(maxRate);
                            partnerName.set(partnerEntry.getKey());
                        }
                    }
                });
            });
            if(min.get() != Integer.MAX_VALUE)
            {
                sb.append("true").append(",").append(partnerName.get()).append(",").append(min.get());
                ((List<InputDetails>)possibleInputs.getList()).add(k);
            }
            else
            {
                sb.append("false").append(",").append("\"\"").append(",").append("\"\"");
                ((List<InputDetails>)imPossibleInputs.getList()).add(k);
            }
            sb.append(System.getProperty("line.separator"));
        });
        System.out.println(sb);
        FileHandling.writeObjectToFile("out1.csv",sb);
    }

    public void getOutput2(List<InputDetails> possibleInputs, Map<String, Partner> partnerList, AtomicInteger totalAmount, List<OutputDetail> result) throws CloneNotSupportedException {
        List<OutputDetail> outputDetail = new ArrayList<>();
        Map<String, Integer> accumulatedOut = new HashMap<>();
        getOutput2Result(possibleInputs, partnerList, outputDetail, accumulatedOut, totalAmount, result);
    }

    private void getOutput2Result(List<InputDetails> possibleInputs, Map<String, Partner> partnerList, List<OutputDetail> outputDetail, Map<String, Integer> accumulatedOut, AtomicInteger totalAmount, List<OutputDetail> result) throws CloneNotSupportedException
    {
        if(possibleInputs.size() == outputDetail.size())
        {
            int amount = getTotalAmount(outputDetail);
            if( amount < totalAmount.get())
            {
                totalAmount.set(amount);
                for(int i=0;i<outputDetail.size();i++)
                {
                    result.add(i, (OutputDetail) outputDetail.get(i).clone());
                }
            }
            return;
        }

        partnerList.entrySet().stream().forEach( partner ->
        {
            if(getPossiblePartner(partner, possibleInputs, outputDetail, accumulatedOut))
            {
                try
                {
                    getOutput2Result(possibleInputs, partnerList, outputDetail, accumulatedOut, totalAmount, result);
                    OutputDetail out = outputDetail.remove(outputDetail.size()-1);
                    int total = accumulatedOut.get(partner.getKey());
                    total-=out.getTotalGB();
                    accumulatedOut.put(partner.getKey(), total);
                } catch (CloneNotSupportedException e) {
                    e.printStackTrace();
                }
            }
        });
    }

    private int getTotalAmount(List<OutputDetail> outputDetail)
    {
        AtomicInteger amount = new AtomicInteger(0);
        outputDetail.stream().forEach(out ->
        {
            amount.addAndGet(out.getTotalRate());
        });
        return amount.get();
    }

    private boolean getPossiblePartner(Map.Entry<String, Partner> partner, List<InputDetails> possibleInputs, List<OutputDetail> outputDetail, Map<String, Integer> accumulatedOut)
    {
        InputDetails inputDetails=possibleInputs.get(outputDetail.size());
        Integer accumulatedGB = accumulatedOut.get(partner.getKey());
        if(accumulatedGB!=null && (accumulatedGB+inputDetails.getNeededGB()) > partner.getValue().getCapacity())
        {
            return false;
        }
        AtomicInteger min = new AtomicInteger(Integer.MAX_VALUE);
        partner.getValue().getTheatreSlabDetails().get(inputDetails.getTheatreName()).stream().forEach(slabDetails ->
        {
            if(slabDetails.getMinSlab() <= inputDetails.getNeededGB() && slabDetails.getMaxSlab() >= inputDetails.getNeededGB())
            {
                int maxRate = slabDetails.getCostPerGB() * inputDetails.getNeededGB();
                if(maxRate < slabDetails.getMinimumRate())
                {
                    maxRate = slabDetails.getMinimumRate();
                }
                if(maxRate < min.get())
                {
                    min.set(maxRate);
                }
            }
        });
        if(min.get() != Integer.MAX_VALUE)
        {
            if (accumulatedGB != null)
            {
                accumulatedOut.put(partner.getKey(), accumulatedGB + inputDetails.getNeededGB());
            } else
            {
                accumulatedOut.put(partner.getKey(), inputDetails.getNeededGB());
            }
            OutputDetail out = new OutputDetail();
            out.setDetails(inputDetails, partner.getKey(), min.get(), inputDetails.getNeededGB());
            outputDetail.add(out);
            return true;
        }
        return false;
    }
}
