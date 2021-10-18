package main

import (
	"fmt"
	"os"
	"tool"
)

func main() {
	// read sample.txt file
	input := os.Args[1]
	inData, errInput := os.ReadFile(input)
	tool.Check(errInput)

	words := tool.SplitWhiteSpaces(string(inData))

	// implement changes for operation tags and 'a'
	words = tool.Operation(words)

	// remove operation tags by only adding elements that are not equal to them to a new string
	str := tool.RemoveTags(words)

	// remove space before punctuation and add space after punctuation by adding to new variable

	str = tool.SortPunctuation(str)

	// change for '''

	str = tool.SortQuotes(str)

	// write to result.txt
	result := []byte(str)
	output := os.Args[2]
	errResult := os.WriteFile(output, result, 0644)
	tool.Check(errResult)

	// read result.txt file
	outData, errOutput := os.ReadFile(output)
	tool.Check(errOutput)

	fmt.Println(string(outData))

}
