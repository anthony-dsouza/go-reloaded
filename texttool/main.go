package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//// read sample.txt file
	dat, err := os.ReadFile(os.Args[1])

	check(err)

	input := string(dat)
	l := len([]rune(input)) - 1
	aRune := []rune(input)

	fmt.Println(input)

	var output string

	var str string

	var op string

	for i, v := range input {
		/// spaces
		if i == l { /// final
			str = str + string(v)
			output = output + str
		} else if v == 40 {
			// find op
			for j := i + 1; aRune[j] != 41; j++ {
				op = op + string(aRune[j])
			}
			fmt.Println(op)
			/// apply op to str then add to output
			if op == "hex" {
				fmt.Println(str)
				/// find last word
				var words []string
				var word string
				for _, v := range str {
					if v != ' ' {
						word = word + string(v)
					} else if v == ' ' {
						words = append(words, word)
						word = ""
					}
				}
				lastindex := len(words) - 1

				lastword := words[lastindex]

				fmt.Println(lastword)

				decimal, err := strconv.ParseInt(lastword, 16, 32)

				dec := strconv.FormatInt(decimal, 10)

				check(err)

				for _, v := range words[:lastindex] {
					output = output + v + " "
				}
				output = output + dec

				fmt.Println(output)

			}
			op = ""
		} else {
			str = str + string(v)
		}
	}
}
