import container.ObjectList;
import handlers.Handler;
import handlers.HandlerFactory;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileOutputStream;
import java.io.FileReader;

public class FileHandling
{
    public static void createObjectFromFile(String fileName, boolean isFirstLineNeeded, ObjectList objectList, Handler handler) {
        if(handler == null)
        {
            System.out.println("Null");
            return;
        }
        try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
            String line = "";
            while ((line = br.readLine()) != null) {
                if (!isFirstLineNeeded) {
                    isFirstLineNeeded = true;
                    continue;
                }
                String[] details = line.split(",");
                handler.handleDetails(details, objectList);
            }

        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void writeObjectToFile(String s, StringBuilder sb)
    {
        try(FileOutputStream fs = new FileOutputStream(new File(s)))
        {
            fs.write(sb.toString().getBytes(),0 , sb.length());
        }catch (Exception e)
        {
            e.printStackTrace();
        }
    }
}
