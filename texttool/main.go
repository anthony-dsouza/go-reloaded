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
	// read sample.txt file
	input := os.Args[1]
	dat, err := os.ReadFile(input)
	Check(err)
	// split by whitespace
	words := tool.SplitWhiteSpaces(string(dat))

	fmt.Println(words)

	for i, word := range words {
		// changes for hex operation
		if word == "(hex)" {
			hexa, err := strconv.ParseInt(words[i-1], 16, 64)
			Check(err)
			words[i-1] = strconv.FormatInt(hexa, 10)
		}
		// changes for bin operation
		if word == "(bin)" {
			bin, err := strconv.ParseInt(words[i-1], 2, 64)
			Check(err)
			words[i-1] = strconv.FormatInt(bin, 10)
		}
		// changes for cap operation
		if word == "(cap)" {
			words[i-1] = strings.Title(words[i-1])
		}
		// changes for cap with respect to number
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
		// changes for low with respect to number
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
		// changes for up with respect to number
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
		// changes for up operation
		if word == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
		}
		// changes for low operation
		if word == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
		}
		// changes for "a"
		if word == "a" {
			// the next element lowered
			lowered := strings.ToLower(words[i+1])
			if lowered[0] == 'h' || lowered[0] == 'a' || lowered[0] == 'e' || lowered[0] == 'i' || lowered[0] == 'o' || lowered[0] == 'u' {
				words[i] = "an"
			}
		}
	}

	// remove operation tags by only adding elements that are not equal to them to a new string
	words1 := ""
	for i, word := range words {
		// removing operation tags followed by numbers
		if word == "(cap," || word == "(up," || word == "(low," {
			// sets them to nothing and when it finds nothing it does not add to new variable
			words[i] = ""
			words[i+1] = ""
		} else if word != "(hex)" && word != "(bin)" && word != "(cap)" && word != "(up)" && word != "(low)" && word != "" {
			// add first word without space
			if i == 0 {
				words1 = words1 + word
			} else {
				words1 = words1 + " " + word
			}
		}
	}

	// remove space before punctuation and add space after punctuation by adding to new variable

	words2 := ""

	for i, letter := range words1 {
		// for when i is the last word -- to stop from going out of index when checking ahead
		if i == len(words1)-1 {
			words2 = words2 + string(letter)
		} else if letter == '.' || letter == ',' || letter == '!' || letter == '?' || letter == ':' || letter == ';' {
			// if the word before is a space -- remove the last element from words2 and add current element and assign to the words2 variable
			if words1[i-1] == ' ' {
				// removes all space at the end of the string words2
				words2 = tool.ReSpace(words2)
				words2 = words2 + string(letter)
			} else {
				// otherwise add current element to words2
				words2 = words2 + string(letter)
			}
			// if the next element is not a space or punctuation -- add a space to words2
			if words1[i+1] != ' ' && words1[i+1] != '.' && words1[i+1] != ',' && words1[i+1] != '!' && words1[i+1] != '?' && words1[i+1] != ':' && words1[i+1] != ';' {
				words2 = words2 + " "
			}
		} else {
			// if not a punctuation or the last letter then add the current element to words2
			words2 = words2 + string(letter)
		}

	}

	// change for '''
	words3 := ""

	// removeSpace is used to check if the space already added to word3 needs to be removed
	// it is set to false to begin with because when we encounter the first ''' the space does not need to removed
	removeSpace := false
	for i, letter := range words2 {
		// we will first search for cases were there is a space before or after the '''
		// if there is no space before or after then the ''' is in the correct position at this point
		if letter == 39 && (words2[i-1] == ' ' || words2[i+2] == ' ') {
			// for the very first ''' removeSpace will be false and will just add the ''' to words3 and set removeSpace to true
			// until it has found another ''' removeSpace will remain true.
			if removeSpace {
				words3 = words3[:len(words3)-1]
				words3 = words3 + string(letter)
				removeSpace = false
			} else {
				words3 = words3 + string(letter)
				removeSpace = true
			}
		} else if i > 1 && words2[i-2] == 39 {
			// the index "i" would need to be greater than 1 because when we check for ''' we need to look
			// 2 indexes back because there would need to be space between that needs to be taken into account
			// if removeSpace will be set to true after the first encounter, we will delete the space by removing the
			// last value in words3 and then adding the next element.
			if removeSpace {
				words3 = words3[:len(words3)-1]
				words3 = words3 + string(letter)
			} else {
				words3 = words3 + string(letter)
			}
		} else {
			words3 = words3 + string(letter)
		}
	}

	// write to result.txt
	d1 := []byte(words3)
	output := os.Args[2]
	err1 := os.WriteFile(output, d1, 0644)
	Check(err1)

	// read sample.txt file
	dat1, err := os.ReadFile(output)
	Check(err)

	fmt.Println()
	fmt.Println(string(dat1))

}
