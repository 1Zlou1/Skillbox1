package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadStudentDetails() (string, int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter student details (name, age, grade):")
	scanner.Scan()
	line := scanner.Text()

	fields := strings.Split(line, ",")
	if len(fields) != 3 {
		fmt.Println("Invalid input")
		return "", 0, 0
	}

	name := strings.TrimSpace(fields[0])
	age := 0
	grade := 0
	_, err := fmt.Sscanf(strings.TrimSpace(fields[1]), "%d", &age)
	if err != nil {
		fmt.Println("Invalid age")
		return "", 0, 0
	}
	_, err = fmt.Sscanf(strings.TrimSpace(fields[2]), "%d", &grade)
	if err != nil {
		fmt.Println("Invalid grade")
		return "", 0, 0
	}

	return name, age, grade
}
