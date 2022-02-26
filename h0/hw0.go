package main

import "fmt"

func main() {
	fmt.Println(FizzBuzz(15))
	fmt.Println(IsPrime(7))
	fmt.Println(IsPalindrom("kerek"))
}

func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 != 0 {
		return "Fizz"
	} else if n%5 == 0 && n%3 != 0 {
		return "Buzz"
	} else {
		return "FizzBuzz"
	}
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsPalindrom(str string) bool {
	var check bool = true
	for i := 0; i < len(str)/2; i++ {
		if string(str[i]) != string(str[len(str)-i-1]) {
			check = false
			break
		}
	}

	return check
}
