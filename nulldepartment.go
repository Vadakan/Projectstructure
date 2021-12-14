package main

type NullDepartment struct {
	Name               string
	NumberofProfessors int
}

func getNullDepartment() *NullDepartment {
	return &NullDepartment{}
}

func (n *NullDepartment) GetNumberOfProfessors() int {
	return 0
}

func (n *NullDepartment) GetName() string {
	return "Null Department"
}
