package main

import (
	"fmt"
	"strconv"
	"strings"
)

const showExpectedResult = false
const showHints = false

func calculate(input1 string, input2 string) float64 {

	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	return float1 + float2
}

func main() {

	fmt.Println("input ur nums")
	var w1 string
	var w2 string
	fmt.Scan(&w1)
	fmt.Scan(&w2)
	result := calculate(w1, w2)
	fmt.Println("Your result is:", result)

}
