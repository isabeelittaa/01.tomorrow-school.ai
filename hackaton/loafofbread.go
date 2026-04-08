package piscine

func LoafOfBread(str string) string {
	result := ""
	word := ""
	count := 0
	skip := 0

	for _, r := range str {
		if skip > 0 {
			skip--
			continue
		}
		if r == ' ' {
			continue
		}
		word += string(r)
		count++
		if count == 5 {
			result += word + " "
			word = ""
			count = 0
			skip = 1
		}
	}

	if len(result) == 0 && len(word) < 5 {
		return "\n"
	}

	if word != "" {
		result += word
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	result += "\n"
	return result
}
