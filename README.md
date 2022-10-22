# Qube Cinemas Challenge 2019
Qube delivers the movie content to theatres all around the world. There are multiple delivery partners to help us deliver the content.

Delivery partners specify the rate of delivery and cost in following manner (All costs are in paise):

Table 1:

| Theatre       | Size Slab (in GB)        | Minimum cost  | Cost Per GB | Partner ID |
| ------------- |:----------------:        |:-------------:| :----------:|:----------:|
|  T1           |         0-200            |       2000    |      20     |     P1     |
|  T1           |         200-400          |       3000    |      15     |P1          |
|  T3           |         100-200          |       4000    |      30     |P1          |
|  T3           |         200-400          |       5000    |      25     |P1          |
|  T5           |         100-200          |       2000    |      30     |P1          |
|  T1           |         0-400            |       1500    |      25     |P2          |

First row allows 0 to 200 GB content to be sent to theatre T1 with the rate 20 paise per GB. However, if total cost comes less than minimum cost, minimum cost (2000 paise) will be charged.

*NOTE*: 
- Multiple partners can deliver to same theatre


- Write programs in any language you want. Feel free to hold the datasets in whatever data structure you want, but try not to use external databases - as far as possible stick to your langauage without bringing in MySQL/Postgres/MongoDB/Redis/Etc.

- We've provided a CSV `partners.csv` with the list of all partners, theatres, content size, minimum cost and cost per GB. Please use the data mentioned there for this program instead of data given in Table 1 and 2. The codes you see in csv may be different from what you see in tables, so please always use the codes in the CSV. This Readme is only an example.

This challenge consist of two problems. Aim to solve atleast Problem Statement 1.

## Problem Statement 1
Given a list of content size and Theatre ID, Find the partner for each delivery where cost of delivery is minimum. If delivery is not possible, mark that delivery impossible.

Use the data given in `partners.csv`.


**Input**: A CSV file `input.csv`. Each row containing delivery ID, size of delivery and theatre ID.

**Expected Output**: A CSV `output.csv`. Each row containing delivery ID, indication if delivery is possible (true/false), selected partner and cost of delivery.

#### Sample Scenarios (Based on above table 1):
**INPUT**:
```
D1, 100, T1
D2, 300, T1
D3, 350, T1
```
**OUTPUT**:
```
D1, true, P1, 2000
D2, true, P1, 4500
D3, true, P1, 5250
```
---
**INPUT**:
```
D1, 70, T1
D2, 300, T1
```
**OUTPUT**:
```
D1, true, P2, 1750
D2, true, P1, 4500
```

---
**INPUT**:
```
D1, 70, T3
D2, 300, T1
```
**OUTPUT**:
```
D1, false, "", "" 
D2, true, P1, 4500
```

## Problem Statement 2

Each partner specifies the **maximum capacity** they can serve, across all their deliveries in following manner:

Table 2:

| Partner ID    | Capacity (in GB) |
| ------------- |:----------------:|
|  P1           |        500       |
|  P2           |        300       |

We have provided `capacities.csv` which contain ID and capacities for each partner.

Given a list of content size and Theatre ID, Assign deliveries to partners in such a way that all deliveries are possible (Higher Priority) and overall cost of delivery is minimum (i.e. First make sure no delivery is impossible and then minimise the sum of cost of all the delivery). If delivery is not possible to a theatre, mark that delivery impossible. Take partner capacity into consideration as well.

Use `partners.csv` and `capacities.csv`.

**Input**: Same as Problem statement 1.

**Expected Output**: Same as Problem statement 1.

#### Sample Scenario (Based on above table 1 and 2):
**INPUT**:
```
D1, 100, T1
D2, 240, T1
D3, 260, T1
```

**OUTPUT**:
```
D1, true, P2, 2500
D2, true, P1, 3600
D3, true, P1, 3900
```

**Explanation**: Only partner P1 and P2 can deliver content to T1. Lowest cost of delivery will be achieved if all three deliveries are given to partner P1 (100\*20+240\*15+260\*15 = 9,500). However, P1 has capacity of 500 GB and total assigned capacity is (100+240+260) 600 GB in this case. Assigning any one of the delivery to P2 will bring the capacity under 500. Assigning the D1, D2 and D3 to P2 is increasing the total cost of delivery by 500 (100\*25+240\*15+260\*15-9500), 2400 (100\*20+240\*25+260*15-9500) and 2600 (100\*20+240\*15+260\*25-9500) respectively. Hence, Assigning D1 to P2.

To submit a solution, fork this repo and send a Pull Request on Github.

For any questions or clarifications, raise an issue on this repo and we'll answer your questions as fast as we can.


Solution - 


also mentioned in solution.txt file separately

How to run
1. Clone project to directory(go workspace recommended)
2. open project directory on terminal, run below command
   go build -> to check if build issues any
   go run main.go -> it'll run the project,with server running on port 8000(Do make sure no other program is running on port 8000)

Run these url in any browser or on any API tester like postman
1. To check the final output - i.e if delivery is possible for the given inputs
   http://localhost:8000/QubeCinema/Isdeliverable
2. To see the partner details
   http://localhost:8000/QubeCinema/PartnerDetails
3. To see input details
   http://localhost:8000/QubeCinema/InputDetails
4. To see partner capacity input details
   http://localhost:8000/QubeCinema/PartnerCapacityDetails


Challenge Understanding

Partners.csv
1.This file contains information about theatre
2.it's allowed size slab in the range(in GB's)
3.cost per GB charged by Partner(partner ID given)
4.partner id for the theatre for the given slab range


Problem 1

Input.csv
1. Delivery ID - unique id of each delivery
2. size of delivery - required content size by each theatre
3. Theatre ID - Theatre to which it'll be delivered

Problem Understanding
Three things to tell
1.If delivery is possible for the given theatre for the required content size
2.Who will deliver it - Partner (Partner ID)
3.What will be the cost of delivery
Observations to arrive at solution
1.Theatre required content size will give us in which rage it's lying, so we can arrive at
decision that which partner can deliver to that theatre which can provide it
2.once we get to know which partner is delivering for that theatre and in which slab it's lying,
then we calculate cost which will be occurring for it
if that cost is less than the minimum cost, then the minimum cost will be returned

Output
1.Delivery ID - the delivery id of the given input delivery ID
2.true/false - whether delivery is possible or not, if true - possible, if false - not possible
3.Cost of delivery
4.Delivery Partner, which will be delivering it


Problem 2

Input.csv
1. Delivery ID - unique id of each delivery
2. size of delivery - required content size by each theatre
3. Theatre ID - Theatre to which it'll be delivered

Problem Understanding
It's in sync with first one and the next level check to problem one
1. Now we know it if delivery is possible for the given theatre and with what cost
2. Now we've to minimize the overall cost of all deliveries
3. Every partner holds the capacity, i.e the maximum it can deliver
4. So while assigning delivery partner's to each theatre with the given delivery id
the things to keep in mind is
a. Partner is left with that much capacity that it can deliver,
b. even if the partner is having capacity to delivery with the minimum charge per GB for that theatre,
we need to check if we're getting overall cost minimum for the all deliveries

Output
1. Delivery ID - unique id of each delivery
2. size of delivery - required content size by each theatre
3. Theatre ID - Theatre to which it'll be delivered
Output will be same except one thing, overall cost will be minimized and all the providing partner's will be delivered under capacity

Note
I tried to solve second problem but not able to come with the exact solution, so mentioned the approach to solve it.
