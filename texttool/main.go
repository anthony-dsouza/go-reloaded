package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"tool"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//// read sample.txt file
	input := os.Args[1]
	dat, err := os.ReadFile(input)
	Check(err)
	/// split by whitespace
	words := tool.SplitWhiteSpaces(string(dat))

	fmt.Println(words)

	for i, word := range words {
		/// changes for hex operation
		if word == "(hex)" {
			hexa, err := strconv.ParseInt(words[i-1], 16, 64)
			Check(err)
			words[i-1] = strconv.FormatInt(hexa, 10)
		}
		/// changes for bin operation
		if word == "(bin)" {
			bin, err := strconv.ParseInt(words[i-1], 2, 64)
			Check(err)
			words[i-1] = strconv.FormatInt(bin, 10)
		}
		/// changes for cap operation
		if word == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
		}
		/// changes for cap with respect to number
		if word == "(cap," {
			// finding the index of the last element
			len := len(words[i+1]) - 1
			// the value of the next word without the parenthesis
			numb := words[i+1][:len]
			// converting numb from a string to an integer
			num, _ := strconv.Atoi(numb)
			// capitalizing the words before the operation tag according to the number
			for j := 1; j <= num; j++ {
				words[i-j] = strings.Title(words[i-j])
			}
		}
		/// changes for low with respect to number
		if word == "(low," {
			// finding the index of the last element
			len := len(words[i+1]) - 1
			// the value of the next word without the parenthesis
			numb := words[i+1][:len]
			// converting numb from a string to an integer
			num, _ := strconv.Atoi(numb)
			// lowercase the words before the operation tag according to the number
			for j := 1; j <= num; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
		}
		/// changes for up with respect to number
		if word == "(up," {
			// finding the index of the last element
			len := len(words[i+1]) - 1
			// the value of the next word without the parenthesis
			numb := words[i+1][:len]
			// converting numb from a string to an integer
			num, _ := strconv.Atoi(numb)
			// uppercase the words before the operation tag according to the number
			for j := 1; j <= num; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
		}
		/// changes for up operation
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
		}
		/// changes for low operation
		if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
		}
	}

	/// remove operation tags by only adding elements that are not equal to them to a new string
	words1 := ""
	for i, word := range words {
		if 
		if word != "(hex)" && word != "(bin)" && word != "(cap)" && word != "(up)" && word != "(low)" {
			// add first word without space
			if i == 0 {
				words1 = words1 + word
			} else {
				words1 = words1 + " " + word
			}

		}
	}

	fmt.Println(words1)

}
