package main

import (
	"fmt"
	"os"
	"tool"
)

func main() {
	// read sample.txt file
	input := os.Args[1]
	dat, err := os.ReadFile(input)
	tool.Check(err)

	words := tool.SplitWhiteSpaces(string(dat))

	// implement changes for operation tags and 'a'
	words = tool.Operation(words)

	// remove operation tags by only adding elements that are not equal to them to a new string
	str := tool.RemoveTags(words)

	// remove space before punctuation and add space after punctuation by adding to new variable

	str = tool.SortPunctuation(str)

	// change for '''

	str = tool.SortQuotes(str)

	// write to result.txt
	d1 := []byte(str)
	output := os.Args[2]
	err1 := os.WriteFile(output, d1, 0644)
	tool.Check(err1)

	// read sample.txt file
	dat1, err := os.ReadFile(output)
	tool.Check(err)

	fmt.Println(string(dat1))

}
