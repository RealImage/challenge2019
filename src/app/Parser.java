package app;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Collections;
import java.util.List;

/**
 * Parser
 */
public abstract class Parser {

    private List<String> getFileData(String filename){

        File file = new File(filename);
        System.out.println(file.getAbsolutePath());
        if(file.exists() && file.canRead()){
            try{
                long start = System.currentTimeMillis(),end = 0;
                List<String> lines = Files.readAllLines(Paths.get(file.getAbsolutePath()));
                end = System.currentTimeMillis();
                System.out.println("Time took " + (end - start) +"ms \n ");
                return lines;
            }catch(IOException e){
                System.err.println("file read error");
            }
        }else{
            System.err.println("file doesn't present or permission issue");
        }

        return Collections.emptyList();
    }

    public final void parseCSV(String file,String separator,boolean header){
        List<String> lines = getFileData(file);
        if(lines.size() > 0 && header){
            lines.remove(0); //remove header
        }
        for (String line : lines) {
            String[] values = line.split(separator);
            mapper(values);
        }
    }

    protected abstract void mapper(String[] values);
    
}