# Challenge2019
Qube Delivers the movie content to theatres all around the world. There are multiple delivery partners to help us deliver the content.

Delivery partners specify the rate of delivery and price in following manner:

| Theatre       | Content Size Slab (in GB)| Minimum Price | Cost Per GB | Partner ID |
| ------------- |:----------------:        |:-------------:| :----------:|:----------:|
|  T1           |         0-200            |       2000    |      20     |     P1     |
|  T1           |         201-300          |       3000    |      15     |P1          |
|  T3           |         101-200          |       4000    |      30     |P1          |
|  T3           |         201-400          |       5000    |      25     |P1          |
|  T5           |         101-200          |       2000    |      30     |P1          |
|  T1           |         0-400            |       1500    |      25     |P2          |

First rows allows 0 to 200 GB content to be sent to theatre T1 with the rate 20 cents per GB. However, if total price comes less than minimum price, minimum price (2000 cents) will be charged.

*NOTE*: Multiple partners can deliver to same theatre 


## Problem Statement
Given a list of content size and Theatre ID to deliver content to, Find the partner where cost of delivery is minimum. If delivery is not possible to a theatre, mark that delivery impossible. 

### Sample Scenarios (Based on above table):
**INPUT**:
```
Number of Deliveries: 1
Content Size of delivery 1: 100
Theatre ID of delivery 1  : T1
Content Size of delivery 2: 300
Theatre ID of delivery 2  : T1
```
**OUTPUT**
```
Solution: 
Delivery: 1, Partner: P1, Cost: 2000
Delivery: 2, Partner: P1, Cost: 4500
Total Cost: 6500
```
---
**INPUT**:
```
Number of Deliveries: 2
Content Size of delivery 1: 70
Theatre ID of delivery 1  : T1
Content Size of delivery 2: 300
Theatre ID of delivery 2  : T1
```
**OUTPUT**
```
Solution: 
Delivery: 1, Partner: P2, Cost: 1750
Delivery: 2, Partner: P1, Cost: 4500
Total Cost: 6250
```

---
**INPUT**:
```
Number of Deliveries: 1
Content Size of delivery 1: 70
Theatre ID of delivery 1  : T3
Content Size of delivery 2: 300
Theatre ID of delivery 2  : T1
```
**OUTPUT**
```
Solution: 
Delivery: 1, Impossible
Delivery: 2, Partner: P1, Cost: 4500
Total Cost: 4500
```
We've provided a CSV with the list of all partners, theatres, content size, minimum price and cost per GB. Please use the data mentioned there for this program. The codes you see there may be different from what you see here, so please always use the codes in the CSV. This Readme is only an example.

Write a program in any language you want that does this. Feel free to make your own input and output format / command line tool / GUI / Webservice / whatever you want. Feel free to hold the dataset in whatever structure you want, but try not to use external databases - as far as possible stick to your langauage without bringing in MySQL/Postgres/MongoDB/Redis/Etc.




