package piscine

func Compact(ptr *[]string) int {
	if ptr == nil || *ptr == nil {
		return 0
	}

	compacted := make([]string, 0, len(*ptr))
	count := 0

	for _, s := range *ptr {
		if s != "" {
			compacted = append(compacted, s)
			count++
		}
	}

	*ptr = compacted

	return count
}
