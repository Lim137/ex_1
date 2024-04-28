package main

import "fmt"

func getNumberWithProperWordForm(value int) string {
	result := ""
	if value < 0 {
		result += "Не бывает -"
		value *= -1
	}
	switch {
	case value%100 > 10 && value%100 < 20:
		result += fmt.Sprintf("%v компьютеров", value)
		return result
	case value%10 == 1:
		result += fmt.Sprintf("%v компьютер", value)
		return result
	case value%10 == 2 || value%10 == 3 || value%10 == 4:
		result += fmt.Sprintf("%v компьютера", value)
		return result
	default:
		result += fmt.Sprintf("%v компьютеров", value)
		return result
	}
}

func main() {
	var value int
	fmt.Scan(&value)
	result := getNumberWithProperWordForm(value)
	fmt.Println(result)
}
