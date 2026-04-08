package piscine

func ToBinary(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		bit := num % 2
		result = string('0'+bit) + result
		num /= 2
	}
	return result
}

func ToDenary(nbr string, base string) int {
	result := 0
	baseLen := len(base)

	for i := 0; i < len(nbr); i++ {
		pos := -1
		for j := 0; j < baseLen; j++ {
			if nbr[i] == base[j] {
				pos = j
				break
			}
		}

		if pos == -1 {
			return -1
		}

		result = result*baseLen + pos
	}
	return result
}

func ToTextBinary(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		digit := num % 2
		result = string('0'+digit) + result
		num = num / 2
	}
	return result
}

func ToTextDecimal(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		digit := num % 10
		result = string('0'+digit) + result
		num = num / 10
	}
	return result
}

func ToTextHexadecimal(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	hexDigits := "0123456789ABCDEF"
	for num > 0 {
		digit := num % 16
		result = string(hexDigits[digit]) + result
		num = num / 16
	}
	return result
}

func ToTextCustomBase(num int, base string) string {
	if num == 0 {
		return string(base[0])
	}
	result := ""
	baseLen := len(base)
	for num > 0 {
		digit := num % baseLen
		result = string(base[digit]) + result
		num = num / baseLen
	}
	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := ToDenary(nbr, baseFrom)
	if num == -1 {
		return "invalid"
	}

	if baseTo == "01" {
		return ToTextBinary(num)
	} else if baseTo == "0123456789" {
		return ToTextDecimal(num)
	} else if baseTo == "0123456789ABCDEF" {
		return ToTextHexadecimal(num)
	} else if len(baseTo) > 1 {
		return ToTextCustomBase(num, baseTo)
	}
	return "unsupported base"
}
