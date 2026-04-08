package piscine

func ConcatParams(args []string) string {
	ln := len(args)
	result := ""
	for i := 0; i < ln; i++ {
		result += args[i]
		if i < ln-1 {
			result += "\n"
		}
	}
	return result
}
