package main

type Student struct {
	studentRollNo int
	studentName   string
	studentAge    int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].studentAge < s[j].studentAge
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
