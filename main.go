package main

import (
	"skillbox/input"
	"skillbox/output"
	"skillbox/student_repo"
)

func main() {
	studentRepo := student_repo.NewMapStudentRepository()

	for {
		name, age, grade := input.ReadStudentDetails()
		if name == "" {
			break
		}

		student := &student_repo.Student{
			Name:  name,
			Age:   age,
			Grade: grade,
		}

		studentRepo.Put(name, student)
	}

	output.PrintStudentNames(studentRepo)
}
