import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
class Solution_problem1{
	
	static ArrayList<Data1> al1=new ArrayList();
	
	public static void Evaluate(String Delivery,int size, String theatre,FileWriter writer) {
		int filtersize=0;
		try
		{		
			BufferedReader csvReader = new BufferedReader(new FileReader("partners.csv"));
			String row1;
			int iteration =0;
	
			while ((row1 = csvReader.readLine()) != null) {
				if(iteration==0) {
					iteration++;
					continue;
				}
				String[] data1 = row1.split(",");
				if((data1[0].trim()).equals(theatre.trim())) {
					filtersize++;
				al1.add(new Data1(data1[0].trim(),data1[1].trim(),data1[2].trim(),data1[3].trim(),data1[4].trim()));
				}
			}
			FinalCalculate(Delivery,size,filtersize,writer);
		}catch(Exception e) {
			e.printStackTrace();
		}
	
		}
	
	private static void FinalCalculate(String Delivery,int size,int filtersize,FileWriter writer) {
		
		int sum=0,minSum=0,count=1;
		String mainPartener="";
		boolean cond=false;
		int handling=0;
		for(int i=0;i<filtersize;i++) {
			
			Data1 a=al1.get(i);
			String p1=a.p1;
			String slab=a.slab;
			String min_cost=a.min_cost;
			String each_cost=a.each_cost;
			String partner=a.partner;
			String[] part = slab.split("-"); 
			
		    int b[] = new int[part.length];
		    for(int j=0;j<b.length;j++){
		      b[j] = Integer.parseInt(part[j]);
		    } 
		    
		    if(b[0]<=size&&b[1]>=size) {
		    	handling++;
		    	cond=true;	
		    	sum=size*(Integer.parseInt(each_cost));
		    	
		    	 if(sum<(Integer.parseInt(min_cost))) {
		    		 
				    	sum=(Integer.parseInt(min_cost));
				    }
		    
		    
		   if(count==1&&sum!=0) {
			   minSum=sum;
			   mainPartener=partner;
			   count++;
		   }else if(sum<=minSum){
			   minSum=sum;
			   mainPartener=partner;
		   }
		   
		    }
		}
		
		if(handling>=1) {
		writeOutput(Delivery,cond,mainPartener,String.valueOf(minSum),writer);
		}else {
		writeOutput(Delivery,cond,"","",writer);
		}
		
	}




	private static void writeOutput(String delivery, boolean cond, String mainPartener, String minSum,FileWriter writer) {
	try {
		String s="";
		if(cond) {
			s="True";
		}else {
			s="False";
		}
		writer.append(delivery);
		writer.append(',');
		writer.append(s);
		writer.append(',');
		writer.append(mainPartener);
		writer.append(',');
		writer.append(minSum);
		writer.append('\n');
		writer.flush();
	}catch(Exception e) {
		System.out.println(e);
	}
	}
	
	public static void main (String args[]) {
try{
	BufferedReader csvReader = new BufferedReader(new FileReader("input.csv"));
	FileWriter writer = new FileWriter("output1.csv");
	
	String row;
	ArrayList<Data> al=new ArrayList();
	
	while ((row = csvReader.readLine()) != null) {
		String[] data = row.split(",");
		al.add(new Data(data[0],Integer.parseInt(data[1]),data[2]));
	    Evaluate(data[0],Integer.parseInt(data[1]),data[2],writer);
	}
}catch(Exception e) {
	System.out.println("the error"+e);
}
}
	}

class Data{
	String a1;
	int a2;
	String a3;
	
	@Override
	public String toString() {
		return a1+a2+a3 ;
	}
	public Data(String a1, int a2, String a3) {
		super();
		this.a1 = a1;
		this.a2 = a2;
		this.a3 = a3;
	}
	
}

class Data1{
	String p1;
	String slab;
	String min_cost;
	String each_cost;
	String partner;
	
	
	@Override
	public String toString() {
		return p1+" "+slab+" "+min_cost+" "+each_cost+" "+partner;
	}
	
	public Data1(String p1, String slab, String min_cost, String each_cost, String partner) {
		super();
		this.p1 = p1;
		this.slab = slab;
		this.min_cost = min_cost;
		this.each_cost = each_cost;
		this.partner = partner;
	}
}