package main

import "fmt"

type Student struct {
	Name  string
	Marks []int
}

func main() {
	s := Student{
		Name:  "Bhavani",
		Marks: []int{23, 45, 46, 78},
	}

	fmt.Printf("Marks: %v\n", s.Marks)

	s.AddMarks(92)

	fmt.Printf("After adding marks: %v\n", s.Marks)

	avg := s.CalculateAverage()
	fmt.Printf("Avg marks: %v\n", avg)
}

func (s *Student) AddMarks(mark int) {
	s.Marks = append(s.Marks, mark)
}

func (s *Student) CalculateAverage() float64 {
	if len(s.Marks) == 0 {
		return 0
	}

	total := 0

	for i := 0; i < len(s.Marks); i++ {
		total += s.Marks[i]
	}

	return float64(total) / float64(len(s.Marks))
}
