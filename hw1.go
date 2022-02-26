package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("GO GO GO")
}

func ParsePhone(phone string) string {
	phone = strings.Replace(phone, " ", "", -1)
	phone = strings.Replace(phone, "-", "", -1)

	return "(" + phone[0:3] + ") " + phone[3:6] + "-" + phone[6:]
}

func Anagram(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if len(s1) == len(s2) {
		for i := 0; i < len(s1); i++ {
			if !strings.Contains(s1, string(s2[i])) {
				return false
			}
		}
	}
	return true
}

func FindEven(e []int) []int {
	for i := 0; i < len(e); i++ {
		if e[i]%2 == 1 {
			e = append(e[:i], e[i+1:]...)
		}
	}

	return e
}

func SliceProduct(e []int) int {
	sum := 0
	for i := 0; i < len(e); i++ {
		sum += e[i]
	}
	return sum
}

func Unique(e []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range e {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func InvertMap(kv map[string]int) map[int]string {
	//for key, value range kv
	r := make(map[int]string)
	for k, v := range kv {
		r[v] = k
	}
	return r
}

func TopCharacters(s string, k int) map[rune]int {
	count := 0
	r := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				count++
			}
		}
		if count > k {
			r[rune(s[i])] = count
		}
		count = 0
	}
	return r
}
