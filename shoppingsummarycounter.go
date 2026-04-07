package piscine

func Fields(s string) []string {
	var result []string
	buffer := ""
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			result = append(result, buffer)
			buffer = ""
		} else {
			buffer += string(r)
		}
	}
	result = append(result, buffer)

	return result
}

func ShoppingSummaryCounter(str string) map[string]int {
	summaryMap := make(map[string]int)
	items := Fields(str)

	for _, item := range items {
		summaryMap[item]++
	}

	return summaryMap
}
