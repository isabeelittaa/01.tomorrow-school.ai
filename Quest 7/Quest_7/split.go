package piscine

func Split(s, sep string) []string {
	lis := []string{}
	text := ""
	i := 0

	for i < len(s) {
		if len(s[i:]) >= len(sep) && s[i:i+len(sep)] == sep {
			if len(text) > 0 {
				lis = append(lis, text)
				text = ""
			}
			i += len(sep)
		} else {
			text += string(s[i])
			i++
		}
	}

	if len(text) > 0 {
		lis = append(lis, text)
	}

	return lis
}
