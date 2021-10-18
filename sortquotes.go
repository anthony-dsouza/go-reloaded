package tool

func SortQuotes(s string) string {
	str := ""

	// removeSpace is used to check if the space already added to word3 needs to be removed
	// it is set to false to begin with because when we encounter the first ''' the space does not need to removed
	removeSpace := false
	for i, letter := range s {
		// we will first search for cases were there is a space before or after the '''
		// if there is no space before or after then the ''' is in the correct position at this point
		if letter == 39 && (s[i-1] == ' ' || s[i+2] == ' ') {
			// for the very first ''' removeSpace will be false and will just add the ''' to str and set removeSpace to true
			// until it has found another ''' removeSpace will remain true.
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(letter)
				removeSpace = false
			} else {
				str = str + string(letter)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 {
			// the index "i" would need to be greater than 1 because when we check for ''' we need to look
			// 2 indexes back because there would need to be space between that needs to be taken into account
			// if removeSpace will be set to true after the first encounter, we will delete the space by removing the
			// last value in str and then adding the next element.
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(letter)
			} else {
				str = str + string(letter)
			}
		} else {
			str = str + string(letter)
		}
	}
	return str
}
