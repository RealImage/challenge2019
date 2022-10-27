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