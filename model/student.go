package model

import (
	"strconv"
)

type Student struct {
	Id      int
	Name    string
	Surname string
}

type StudentDto struct {
	Name    string
	Surname string
}

func (s *StudentDto) ToString() string {
	return s.Name + " " + s.Surname
}

func ToDto(student Student) StudentDto {
	return StudentDto{
		Name:    student.Name,
		Surname: student.Surname,
	}
}

func Initialize(count int) []Student {

	var students []Student

	for i := 0; i < count; i++ {
		strI := strconv.Itoa(i)
		student := Student{
			Id: i,
			Name:    "n" + strI,
			Surname: "s" + strI,
		}
		students = append(students, student)
	}

	return students
}