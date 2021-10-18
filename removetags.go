package tool

func RemoveTags(s []string) string {
	str := ""
	for i, word := range s {
		// removing operation tags followed by numbers
		if word == "(cap," || word == "(up," || word == "(low," {
			// sets next element which will be number to nothing so that it is not added
			s[i+1] = ""
		} else if word != "(hex)" && word != "(bin)" && word != "(cap)" && word != "(up)" && word != "(low)" && word != "" {
			// add first word without space
			if i == 0 {
				str = str + word
			} else {
				str = str + " " + word
			}
		}
	}
	return str
}
