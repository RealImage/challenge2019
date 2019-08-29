## RealImage challenge 2019
[Problem statement Link](https://github.com/RealImage/challenge2019/blob/master/README.md)
### Tech stack
- golang v1.10.4

### Execution
- The extracted Qube folder should reside within go path.
- `cd src/qube` and `./bin/main`
<br> or <br>
- `cd src/qube` and `go run main`

### Folder structure
- module based structure. `src` folder will contain modules.
- main.go file resides in `module` root 
- `module/static` contains input and output(generated automatically after running the code) csv.
- `module/structs` contains structs
- `module/prepocess` contains file which converts and stores 
    the data read from input csv file in data structure.
 
### Misc
- Refer `main.go` for solution for both the problems.
- For problem 2 optimal solution would be using `min-cost max-flow` algorithm. Since I do not have required knowledge 
of the algorithm, have implement another logic which is explained in comments section in respective file.
The applied logic is not optimal but should work for most of the use cases.