package main

type computerscience struct {
	Name               string
	NumberOfProfessors int
}

func getComputerscience(Name string, nump int) *computerscience {
	return &computerscience{
		Name:               Name,
		NumberOfProfessors: nump,
	}
}

func (cs *computerscience) GetNumberOfProfessors() int {
	return cs.NumberOfProfessors
}

func (cs *computerscience) GetName() string {
	return cs.Name
}
