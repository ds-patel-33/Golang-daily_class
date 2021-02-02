package main

import (
	"fmt"
)

func main() {

	//initialize map without data
	m := make(map[string]int)
	m["route1"] = 66
	m["route2"] = 68
	fmt.Println("map m:", m)

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	//initialize map with some data
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	fmt.Println("map commits:", commits)
	fmt.Println("length of commits:", len(commits))

	//checking the value exists  or not
	var employees = map[int]string{
		1001: "John",
		1002: "Steve",
		1003: "Maria",
	}

	printEmployee(employees, 1001)
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}

	//deleting map
	var fileExtensions = map[string]string{
		"Python": ".py",
		"C++":    ".cpp",
		"Java":   ".java",
		"Golang": ".go",
		"Kotlin": ".kt",
	}

	fmt.Println(fileExtensions)

	delete(fileExtensions, "Kotlin")
	delete(fileExtensions, "Javascript")

	fmt.Println(fileExtensions)

}

func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}
