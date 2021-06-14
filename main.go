package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// An Employee is a record of an employee name, Salary, and their Manager.
type Employee struct {
	Name         string      `json:"name"`
	Salary       float64     `json:"salary"`
	Manager_Name string      `json:"manager_name"` // functions as a Foreign Key
	Manager      *Employee   // Pointer to a manager Employee
	Manages      []*Employee // Slice of Employee pointers
}

// A map[string] of type Employee to unmarshal json to
// If there were an outer object, we would need to export this
// https://stackoverflow.com/a/32751792 -> map value as a pointer
type employees struct {
	data map[string]*Employee
}

// Acts as a constructor method to the in-memory database
func newEmployees() *employees {
	return &employees{
		data: map[string]*Employee{},
	}
}

// Accepts Employee name to be used as Primary Key for database
// In another scenario, we could implement an ID field to use as a PK
func (e *employees) getEmployee(emp string) (*Employee, error) {
	employee, ok := e.data[emp]
	if !ok {
		err := fmt.Errorf("employee record: %s not found", emp)
		return nil, err
	}
	return employee, nil
}

// Accepts two strings: employee name and manager name to use as PK's in the db
func (e *employees) setManager(emp_name, manager_name string) error {
	// Ensure that employee exists
	_, err := e.getEmployee(emp_name)
	if err != nil {
		return err
	}

	// Ensure Manager Employee exists, return the pointer to correct Employee
	managerPtr, err := e.getEmployee(manager_name)
	if err != nil {
		return err
	}

	// Set the Manager field to the pointer to the Employee
	e.data[emp_name].Manager = managerPtr
	return nil
}

// Each Employee has a Manages slice, pointers to other Employees that they manage
func (e *employees) addToManagesList(employee, manager string) error {
	employeePtr, err := e.getEmployee(employee)
	if err != nil {
		return err
	}

	_, err = e.getEmployee(manager)
	if err != nil {
		return err
	}

	// Append the Manages field slice with the pointer to the Employee
	e.data[manager].Manages = append(e.data[manager].Manages, employeePtr)
	return nil
}

// Sets up the relationships from the JSON file
func (e *employees) setRelations() {
	for k, v := range e.data {
		// Compare to zero-value of struct field, if doesn't exist -> skip
		if v.Manager_Name != "" {
			e.setManager(k, v.Manager_Name)
			e.addToManagesList(k, v.Manager_Name)
		}
	}
}

// Assuming we don't know anything about Employee structure, we need to find
// an Employee that has no manager to find the top-level
func (e *employees) findRootEmployee() *Employee {
	for _, employee := range e.data {
		if employee.Manager == nil {
			return employee
		}
	}
	return nil
}

// Recursively prints an Employee that manages and prints all Employee that reports to them
func (e *employees) printManager(manager *Employee) {
	// fmt.Printf(manager.Name + "\n")
	// If an Employee manages two or more Employee entities, we sort before printing
	if len(manager.Manages) > 0 {
		fmt.Printf("Employees of: %s\n", manager.Name)
		// Sort the Employee entities that are managed, by name
		sort.SliceStable(manager.Manages, func(i, j int) bool {
			return manager.Manages[i].Name < manager.Manages[j].Name
		})
	}

	// Loop through all the Employee entities that are managed
	for _, employee := range manager.Manages {
		fmt.Printf("\t")
		fmt.Printf(employee.Name + "\n")
	}

	// Check for another manager amongst managed Employee entities
	for _, employee := range manager.Manages {
		if employee.Manages != nil {
			e.printManager(employee)
		}
	}
}

// Finds the first Employee to print and begins call to recursive function
func (e *employees) printEmployeeStructure() error {
	// Find the top-level Employee
	manager := e.findRootEmployee()
	if manager == nil {
		return fmt.Errorf("cannot find top-level employee")
	}
	// Print the top-level Employee name
	fmt.Printf(manager.Name + "\n")
	e.printManager(manager)
	return nil
}

// Returns a summation of total company salary expenses
func (e *employees) sumSalary() float64 {
	sum := 0.0
	for _, v := range e.data {
		sum += v.Salary
	}
	return sum
}

func ReadFile(filename string) (*employees, error) {
	// Open and read the JSON file
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize the in-memory data store
	e := newEmployees()

	// Unmarshal the JSON to the employees structure
	json.Unmarshal(byteValue, &e.data)
	return e, nil
}

func main() {
	employees, err := ReadFile("./employees.json")
	if err != nil {
		panic("error reading json file")
	}

	// Set relationships between Employee
	employees.setRelations()

	// Print the Employees
	err = employees.printEmployeeStructure()
	if err != nil {
		fmt.Println(err)
	}

	// Print the Total salary expense of company
	fmt.Printf("Total salary: %.2f\n", employees.sumSalary())
}
