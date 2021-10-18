package tool

func SortPunctuation(s string) string {
	str := ""

	for i, letter := range s {
		// for when i is the last word -- to stop from going out of index when checking ahead
		if i == len(s)-1 {
			if letter == '.' || letter == ',' || letter == '!' || letter == '?' || letter == ':' || letter == ';' {
				// if the word before is a space -- remove the last element from str and add current element and assign to the str variable
				if s[i-1] == ' ' {
					// removes all space at the end of the string str
					str = ReSpace(str)
					str = str + string(letter)
				} else {
					// otherwise add current element to str
					str = str + string(letter)
				}
			} else {
				str = str + string(letter)
			}
		} else if letter == '.' || letter == ',' || letter == '!' || letter == '?' || letter == ':' || letter == ';' {
			if s[i-1] == ' ' {
				str = ReSpace(str)
				str = str + string(letter)
			} else {
				str = str + string(letter)
			}
			// if the next element is not a space or punctuation -- add a space to str
			if s[i+1] != ' ' && s[i+1] != '.' && s[i+1] != ',' && s[i+1] != '!' && s[i+1] != '?' && s[i+1] != ':' && s[i+1] != ';' {
				str = str + " "
			}
		} else {
			// if not a punctuation or the last letter then add the current element to str
			str = str + string(letter)
		}

	}
	return str
}
