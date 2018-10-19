package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Printf("Fizz ")
		} else if i%5 == 0 {
			fmt.Printf("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}

func fizzBuzz2() {
	for i := 1; i <= 100; i++ {
		s := ""
		if i%3 == 0 || containsDigit(i, 3) {
			s += "Fizz"
		}
		if i%5 == 0 || containsDigit(i, 5) {
			s += "Buzz"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		fmt.Printf("%s ", s)
	}
}

func containsDigit(n, target int) bool {
	targetDigit := strconv.Itoa(target)
	digits := strings.Split(strconv.Itoa(n), "")
	for _, digit := range digits {
		if digit == targetDigit {
			return true
		}
	}
	return false
}

func main() {
	// fizzBuzz()
	fizzBuzz2()
}
