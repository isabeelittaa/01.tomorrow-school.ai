package piscine

func Rot14(s string) string {
	text := ""

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = c + 14
			if c > 'z' {
				c = c - 26
			}
			text += string(c)
		} else if c >= 'A' && c <= 'Z' {
			c = c + 14
			if c > 'Z' {
				c = c - 26
			}
			text += string(c)
		} else {
			text += string(c)
		}
	}

	return text
}
