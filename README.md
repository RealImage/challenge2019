# Qube Cinemas Challenge 2019 - Solution by Shreeyash Naik

## Solution for Problem Statement 1

**INPUT**:
```
D1,150,T1
D2,325,T2
D3,510,T1
D4,700,T2
```
**MY OUTPUT**:
```
D1,true ,P1,2000
D2,true ,P1,3250
D3,true ,P3,15300
D4,false,,
```

*ALGORITHM*:
1. Stores details from partner.csv as a mapping of theatre_id -> all partner details.
2. For theatre_id from a delivery_id, iterates over the partner_details and computes the minimum cost. 


---
<br/>

## Solution for Problem Statement 2

**INPUT**:
```
D1,150,T1
D2,325,T2
D3,510,T1
D4,700,T2
```

**MY OUTPUT**:
```
D1,true,P1,2000
D2,true,P2,3500
D3,true,P3,15300
D4,false,,
```
**NOTE**: 
For problem statement 2, my output is different than expected output. However, the solution satisfies all the conditions, plus overall cost of delivery (2000+3500+15300) is lesser than that of expected output.

*ALGORITHM*:
1. Stores details from partner.csv as a mapping of theatre_id -> all partner details.
2. Stores in array, for each delivery_id in input.csv stores size_requirement, number of valid available Partner Options, details of each partner options.
3. Sort the above array on the basis number of available options.
4. Delivery with least number of options will get processed first.


<br />

## Tree Layout

```tree
.

├── go.mod
├── main.go
├── README.md
├── common
│   ├── db
│   │   ├── db.go
│   │   └── utils.go
│   ├── schemas
│   │   ├── schemas.go
│   ├── utils
│   │   ├── fileutils.go
│   │   └── utils.go
├── config
│   └── config.go
├── README.md
└── src
│   │── compute.go
│   └── solutions.go

```