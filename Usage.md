
The `output2.csv` is wrong and it should contain the following data
```
D1,true ,P1,2000
D2,true ,P2,3500
D3,true ,P3,15300
D4,false,"",""
```
As this the case with total cost minimum and with maximum delivery.


To get solution of problem1 specify partners,input and path of output file

For problem2 specify the capacities file along with the above files



Examples:

For Problem-1:

```
go build -o ./bin/main && ./bin/main -inputFile input/input.csv -partnersFile input/partners.csv -outputFile output/output3.csv


go build -o ./bin/main && ./bin/main -inputFile input/input2.csv -partnersFile input/partners2.csv -outputFile output/output3.csv

go build -o ./bin/main && ./bin/main -inputFile input/input3.csv -partnersFile input/partners2.csv -outputFile output/output3.csv


go build -o ./bin/main && ./bin/main -inputFile input/input4.csv -partnersFile input/partners2.csv -outputFile output/output3.csv
```

For Problem-2:

```
go build -o ./bin/main && ./bin/main -inputFile input/input5.csv -partnersFile input/partners2.csv -outputFile output/output5.csv -capacitiesFile input/capacities1.csv


go build -o ./bin/main && ./bin/main -inputFile input/input.csv -partnersFile input/partners.csv -outputFile output/output4.csv -capacitiesFile input/capacities.csv
```