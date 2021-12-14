package main

import "fmt"

func main() {

	RMKCollege := Initcollege()
	RMKCollege.Adddepartment("Mechanical")
	RMKCollege.Adddepartment("CSE")

	obj := RMKCollege.getDepartment("Mechanical")

	fmt.Println(obj.GetName())
	fmt.Println(obj.GetNumberOfProfessors())

	obj = RMKCollege.getDepartment("EEE")
	fmt.Println(obj.GetName())
	fmt.Println(obj.GetNumberOfProfessors())

}
