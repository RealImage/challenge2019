package container;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public class InputList implements ObjectList
{
    List<InputDetails> inputs = new ArrayList<>();
    @Override
    public List getList() {
        return inputs;
    }
}
