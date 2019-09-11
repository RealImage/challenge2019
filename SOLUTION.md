## RealImage challenge 2019
[Problem statement Link](https://github.com/RealImage/challenge2019/blob/master/README.md)
### Tech stack
- golang v1.10.4
- Also tested in v1.12.9
### Execution
- Import the package using `go get -v github.com/funcoding/challenge2019`
- To execute the package run `./${GOPATH}/bin/challenge2019`

### Folder structure
- main.go file resides in `root 
- `static` contains input and output(generated automatically after running the code) csv.
- `structs` contains structs
- `prepocess` contains file which converts and stores 
    the data read from input csv file in data structure.
 
### Misc
- Refer `main.go` for solution for both the problems.
- For problem 2 optimal solution would be using `min-cost max-flow` algorithm. Since I do not have required knowledge 
of the algorithm, have implement another logic which is explained in comments section in respective file.
The applied logic is not optimal but should work for most of the use cases.