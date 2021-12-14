package main

type college struct {
	department []Department
}

func Initcollege() *college {
	return &college{}
}

func (c *college) Adddepartment(name string) {
	if name == "Mechanical" {
		c.department = append(c.department, getMech(10, name))
	} else if name == "CSE" {
		c.department = append(c.department, getComputerscience(name, 20))
	}

}

func (c *college) getDepartment(name string) Department {
	for _, department := range c.department {
		if name == department.GetName() {
			return department
		}
	}
	return getNullDepartment()
}
