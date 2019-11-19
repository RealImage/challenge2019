using System;
using System.Collections.Generic;
using System.Text;
using System.IO;

namespace Qubechallenge
{
    class Program
    {
        static int header = 0,iterationcount=0;
        static List<Partners> Partnerlist = new List<Partners>();
         
       
        public static void Main(string[] args)
        {
            List<input> inputarray = new List<input>();
            var inputReader = new StreamReader(File.OpenRead(@"C:\Users\Suryakumar\Desktop\Project\Qubechallenge\Data\input.csv"));
            var capacityReader = new StreamReader(File.OpenRead(@"C:\Users\Suryakumar\Desktop\Project\Qubechallenge\Data\capacities.csv"));

            var map = new Dictionary<string, int>(); // Dictionary to store capacity of each partner

            int linecounter = 0;

            while (!capacityReader.EndOfStream)
            {
                if(linecounter==0) //To skip header line
                {
                    var Line = capacityReader.ReadLine();
                    linecounter++;
                }

                if(linecounter>0) // Access the capacities given
                {
                    var Line = capacityReader.ReadLine();
                    var row = Line.Split(',');
                    map.Add(row[0].TrimEnd(), Convert.ToInt32(row[1]));
                   
                }
                
            }

            int pblm=1;

            //Getting User Input

            Console.WriteLine("Enter the Problem number :");
            Console.WriteLine("1.Problem Statement-1 (Press 1)");
            Console.WriteLine("1.Problem Statement-2 (Press 2");

            pblm = Convert.ToInt32(Console.ReadLine());

            //Iterate for each input given and call the respective methods

            while (!inputReader.EndOfStream)
            {
                var Line = inputReader.ReadLine();
                var row = Line.Split(',');

                
                input ip = new input()
                {
                    deliveryId=row[0],
                    deliverySize=Convert.ToInt32(row[1]),
                    theatreId=row[2]
                };
                inputarray.Add(ip);

                LoadPartnersList(row[0], Convert.ToInt32(row[1]), row[2],map,pblm);

            }
        }

        //Method to load the relevant partners who can deliver to the current theatreid into the Partners list
        public static void LoadPartnersList(string deliveryid, int deliverysize, string theatreid, Dictionary<string, int> map, int problemNo)
        {
            var Partnerreader = new StreamReader(File.OpenRead(@"C:\Users\Suryakumar\Desktop\Project\Qubechallenge\Data\partners.csv"));
            Partnerlist.Clear();

            while (!Partnerreader.EndOfStream)
            {
                if (header == 0)       //To skip the header row
                {
                    var Line = Partnerreader.ReadLine();
                    header++;
                    continue;
                }

                if (header > 0)        //Logic to load the list
                {
                    var Line = Partnerreader.ReadLine();
                    var row = Line.Split(',');
                    var minmax = row[1].Split('-');
                    if (row[0].TrimEnd().Equals(theatreid))
                    {
                        iterationcount++;
                        Partners P = new Partners()
                        {
                            theatreId = row[0],
                            min = Convert.ToInt32(minmax[0]),
                            max = Convert.ToInt32(minmax[1]),
                            minCost = Convert.ToInt32(row[2]),
                            costPerGb = Convert.ToInt32(row[3]),
                            partnerId = row[4]
                        };
                        Partnerlist.Add(P);
                    }
                }

            }

            if (problemNo == 1)
            {
                CheckDelivery(deliveryid, deliverysize, iterationcount, theatreid);
            }

            if (problemNo == 2)
            {
                CheckDeliverywithCapacity(Partnerlist, deliveryid, deliverysize, theatreid, map);
            }

        }

        //Method to check the possibility of delivery and amount calculation
        public static void CheckDelivery(string deliveryid, int deliverysize, int iterationcount, string theatreid)
        {

            int amount = 0, curramount = int.MaxValue;
            bool deliveryCondition = false;
            string partner = "";
                        
            for (int i = 0; i < Partnerlist.Count; i++)
            {
                string theatreIdObj = Partnerlist[i].theatreId;
                int minObj = Partnerlist[i].min;
                int maxObj = Partnerlist[i].max;
                int minCostObj = Partnerlist[i].minCost;
                int costPerGbObj = Partnerlist[i].costPerGb;
                string partnerIdobj = Partnerlist[i].partnerId;

                if (deliverysize > minObj && deliverysize < maxObj)
                {
                    deliveryCondition = true;
                    amount = (deliverysize * costPerGbObj);

                    if (amount < minCostObj)
                    {
                        amount = minCostObj;
                    }

                    if (amount < curramount)
                    {
                        curramount = amount;
                        partner = partnerIdobj;
                    }
                }

            }

            writeOutput(deliveryid, deliveryCondition, partner, curramount);
        }


        //Method to check the delivery based on Capacities given
        public static void CheckDeliverywithCapacity(List<Partners> T, string deliveryid, int deliverysize, string theatreid, Dictionary<string, int> map)
        {

            int amount = 0, curramount = int.MaxValue;
            bool condition = false;
            string partner = "";
            
            for (int i = 0; i < Partnerlist.Count; i++)
            {
                string theatreIdObj = Partnerlist[i].theatreId;
                int minObj = Partnerlist[i].min;
                int maxObj = Partnerlist[i].max;
                int minCostObj = Partnerlist[i].minCost;
                int costPerGbObj = Partnerlist[i].costPerGb;
                string partnerIdobj = Partnerlist[i].partnerId;

                int capacity = map[partnerIdobj];

                if (deliverysize > minObj && deliverysize < maxObj && deliverysize < capacity)
                {
                    condition = true;
                    amount = (deliverysize * costPerGbObj);

                    if (amount < minCostObj)
                    {
                        amount = minCostObj;
                    }

                    if (amount < curramount)
                    {
                        curramount = amount;
                        partner = partnerIdobj;
                        map[partnerIdobj] = capacity - deliverysize;
                    }
                }

            }


            writeOutput(deliveryid, condition, partner, curramount);

        }

        public static void writeOutput(string deliveryId, bool deliveryCondition, string partnerId, int amount)
        {
            var csvwriter = new StringBuilder();
            var newLine = "";
            var filepath = @"C:\Users\Suryakumar\Desktop\Project\Python Projects\Qubechallenge\Qubechallenge\Data\output2.csv";
            string writeAmount=amount.ToString();

            if(deliveryCondition==false)
            {
                partnerId = "";
                writeAmount = "";
            }

            newLine = string.Format("{0},{1},{2},{3}", deliveryId, deliveryCondition, partnerId, writeAmount);

            csvwriter.AppendLine(newLine);
            File.AppendAllText(filepath, csvwriter.ToString());

        }
    }

    //Getter and Setter methods
    public class Partners
    {
        public string theatreId { get; set; }
        public int min { get; set; }
        public int max { get; set; }
        public int minCost { get; set; }
        public int costPerGb { get; set; }
        public string partnerId { get; set; }

    }

    public class input
    {
        public string deliveryId { get; set; }
        public int deliverySize { get; set; }
        public string theatreId { get; set; }
        
    }

}
