package main

type Mech struct {
	NumberOfProfessors int
	Name               string
}

func getMech(NumberOfProfessors int, Name string) *Mech {
	return &Mech{
		NumberOfProfessors: NumberOfProfessors,
		Name:               Name,
	}
}
