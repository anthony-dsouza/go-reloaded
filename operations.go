package tool

import (
	"strconv"
	"strings"
)

func Operation(s []string) []string {
	for i, word := range s {
		// changes for hex operation
		if word == "(hex)" {
			hexa, err := strconv.ParseInt(s[i-1], 16, 64)
			Check(err)
			s[i-1] = strconv.FormatInt(hexa, 10)
		}
		// changes for bin operation
		if word == "(bin)" {
			bin, err := strconv.ParseInt(s[i-1], 2, 64)
			Check(err)
			s[i-1] = strconv.FormatInt(bin, 10)
		}
		// changes for cap operation
		if word == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
		}
		// changes for cap with respect to number
		if word == "(cap," {
			// finding the index of the last element
			len := len(s[i+1]) - 1
			// the value of the next word without the parenthesis
			numb := s[i+1][:len]
			num, numerr := strconv.Atoi(numb)
			Check(numerr)
			// capitalizing the words before the operation tag according to the number
			for j := 1; j <= num; j++ {
				s[i-j] = strings.Title(s[i-j])
			}
		}
		// changes for low with respect to number
		if word == "(low," {
			len := len(s[i+1]) - 1
			numb := s[i+1][:len]
			num, _ := strconv.Atoi(numb)
			// lowercase the words before the operation tag according to the number
			for j := 1; j <= num; j++ {
				s[i-j] = strings.ToLower(s[i-j])
			}
		}
		// changes for up with respect to number
		if word == "(up," {
			len := len(s[i+1]) - 1
			numb := s[i+1][:len]
			num, _ := strconv.Atoi(numb)
			// uppercase the words before the operation tag according to the number
			for j := 1; j <= num; j++ {
				s[i-j] = strings.ToUpper(s[i-j])
			}
		}
		// changes for up operation
		if word == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
		}
		// changes for low operation
		if word == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
		}
		// changes for "a"
		if word == "a" {
			// the next element lowered and assigned to variable to reduce letter comparisons
			lowered := strings.ToLower(s[i+1])
			if lowered[0] == 'h' || lowered[0] == 'a' || lowered[0] == 'e' || lowered[0] == 'i' || lowered[0] == 'o' || lowered[0] == 'u' {
				s[i] = "an"
			}
		}
	}

	return s
}
