package main

type NullDepartment struct {
	Name               string
	NumberofProfessors int
}

func getNullDepartment(Name string, NumberOfProf int) *NullDepartment {
	return &NullDepartment{
		Name:               Name,
		NumberofProfessors: NumberOfProf,
	}
}

func (n *NullDepartment) GetNumberOfProfessors() int {
	return 0
}

func (n *NullDepartment) GetName() string {
	return "no name"
}
