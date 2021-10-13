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
	}
	fmt.Println(words)

	/// remove operation tags into new variable
	words1 := ""
	for _, word := range words {
		if word != "(hex)" && word != "(bin)" {
			words1 = words1 + " " + word
		}
	}

	fmt.Println(words1)

}
