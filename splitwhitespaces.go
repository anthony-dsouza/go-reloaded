package tool

func SplitWhiteSpaces(s string) []string {
	var str []string

	l := len(s) - 1

	var word string

	for i, v := range s {
		if i == l {
			word = word + string(v)
			str = append(str, word)
		} else if v == 32 || v == 15 {
			// check for space and tabs
			if s[i+1] == ' ' || s[i+1] == '	' {
			} else {
				str = append(str, word)
				word = ""
			}
		} else {
			word = word + string(v)
		}
	}

	return str
}
