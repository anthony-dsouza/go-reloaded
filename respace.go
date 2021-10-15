package tool

// removes spaces at the end of the string
func ReSpace(s string) string {
	// last index of string
	len := len(s) - 1

	// check to see if there is another space at the end -- if this is true it will apply the function again on the string without the current space
	if s[len-1] == ' ' {
		return ReSpace(s[:len])
	}

	// if no spaces before the current space then it will return the string with no spaces
	return s[:len]
}
