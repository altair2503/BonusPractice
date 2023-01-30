package studentModule

import "strconv"

type Student struct {
	Surname     string
	Name        string
	Age         int
	Department  string
	YearOfStudy int
}

func (st Student) GetName() string {
	return st.Name
}

func (st Student) SetName(name string) {
	st.Name = name
}

func (st Student) GetInfo() string {
	s := st.Name + " " + st.Surname + ":" + " " + strconv.Itoa(st.Age) + " year, departmnet: " + st.Department + ", year of study: " + strconv.Itoa(st.YearOfStudy)
	return s
}
