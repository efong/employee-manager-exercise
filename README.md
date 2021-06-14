# Coding Exercise: Employee-Manager Office Hierarchy
For the following office hierarchy, provide concrete implementations of Manager and Employee. Write a test program which does two things:
- Print out an ASCII employee tree (any format you want)
- Print out the total salary requirements for the entire company.

## Expectations
- Sort employees alphabetically
- Build unit tests for the code
- Create the hierarchy by reading in from a properties file or json/yaml file

## Sample Output
```
Jeff
Employees of: Jeff
    Dave
Employees of: Dave
    Andy
    Dan
    Jason
    Rick
    Suzanne
Total salary: 275000
```

## Expected Input
A JSON file of the following format:
```json
{
  "Jeff": { "name": "Jeff", "salary": 200000 },
  "Dave": { "name": "Dave", "salary": 150000, "manager_name": "Jeff" },
  "Andy": { "name": "Andy", "salary": 100000, "manager_name": "Dave" },
  "Dan": { "name": "Dan", "salary": 100000, "manager_name": "Dave" },
  "Jason": { "name": "Jason", "salary": 100000, "manager_name": "Dave" },
  "Rick": { "name": "Rick", "salary": 100000, "manager_name": "Dave" },
  "Suzanne": { "name": "Suzanne", "salary": 100000, "manager_name": "Dave" }
}
```

## Actual Output
```
Jeff
Employees of: Jeff
        Dave
Employees of: Dave
        Andy
        Dan
        Jason
        Rick
        Suzanne
Total salary: 850000.00
```

### Output with multi-level hierarchy
```
Jeff
Employees of: Jeff
        Dave
Employees of: Dave
        Andy
        Dan
Employees of: Dan
        Jason
        Rick
Employees of: Rick
        Suzanne
Total salary: 850000.00
```
Caveat(s):
- Did not add tests to detect cycles.

## Getting started and testing:

You can test this by running:
```
go build -o testapp
go test -count=1
```

Or simply:

```
make
```
