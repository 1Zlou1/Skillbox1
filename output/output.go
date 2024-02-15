package output

import (
	"fmt"
	"skillbox/student_repo"
)

func PrintStudentNames(repo student_repo.StudentRepository) {
	students := repo.GetStudents()
	for _, student := range students {
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("Grade:", student.Grade)
		fmt.Println()
	}

}
