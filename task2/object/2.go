package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID int
}

func (employee *Employee) printEmployeeInfo() {
	fmt.Printf("EmployeeId：%d，Name: %s, Age: %d\n", employee.EmployeeID, employee.Person.Name, employee.Person.Age)
}

func main() {

	var e = &Employee{
		Person: Person{
			Name: "zhangsan",
			Age:  18,
		},
		EmployeeID: 1,
	}

	e.printEmployeeInfo()
}
