package year

import "fmt"

type Day struct {
	Number int
	Run    func()
}

type Year struct {
	Number int
	Days   []Day
}

func (y *Year) RunRange(start int, end int) {
	y.PrintSeparator()
	for i := start; i <= end; i++ {
		y.Days[i-1].Run()
	}
}

func (y *Year) RunDay(day int) {
	y.PrintSeparator()
	y.Days[day-1].Run()
}

func (y *Year) RunAllDays() {
	y.PrintSeparator()
	for _, day := range y.Days {
		day.Run()
	}
}

func (y *Year) PrintSeparator() {
	fmt.Printf("===== Year %d =====\n", y.Number)
}