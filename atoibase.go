package piscine

func AtoiBase(s string, base string) int {
	// Проверка валидности базы
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, r := range s {
		index := indexInBase(r, base)
		if index == -1 {
			return 0
		}
		result = result*baseLen + index
	}
	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i, a := range base {
		if a == '+' || a == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if a == rune(base[j]) {
				return false
			}
		}
	}
	return true
}

func indexInBase(r rune, base string) int {
	for i, b := range base {
		if b == r {
			return i
		}
	}
	return -1
}
