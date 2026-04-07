package piscine

func SplitWhiteSpaces(s string) []string {
	lis := []string{}
	text := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if len(text) > 0 {
				lis = append(lis, text)
				text = ""
			}
		} else {
			text += string(s[i])
		}
	}
	if len(text) > 0 {
		lis = append(lis, text)
	}
	return lis
}
