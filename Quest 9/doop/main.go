package main

import (
	"os"
)

const (
	MaxInt64 = 9223372036854775807
	MinInt64 = -9223372036854775808
)

func itoa(n int64) []byte {
	if n == 0 {
		return []byte{'0'}
	}
	if n == MinInt64 {
		return []byte("-9223372036854775808")
	}

	isNegative := n < 0
	if isNegative {
		n = -n
	}

	var buf [20]byte
	i := len(buf) - 1

	for n > 0 {
		buf[i] = byte('0' + (n % 10))
		n /= 10
		i--
	}

	if isNegative {
		buf[i] = '-'
		i--
	}

	return buf[i+1:]
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
	os.Stdout.Write([]byte{'\n'})
}

func myAtoi64(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}

	sign := int64(1)
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if len(s) == start {
		return 0, false
	}

	var result uint64 = 0
	for i := start; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			return 0, false
		}

		digit := uint64(char - '0')

		if result > uint64(MaxInt64)/10 {
			return 0, false
		}
		result = result*10 + digit

		if sign == 1 {
			if result > uint64(MaxInt64) {
				return 0, false
			}
		} else {
			if result > uint64(MaxInt64)+1 {
				return 0, false
			}
		}
	}

	if sign == -1 {
		return -int64(result), true
	}
	return int64(result), true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	val1Str := os.Args[1]
	op := os.Args[2]
	val2Str := os.Args[3]

	val1, ok1 := myAtoi64(val1Str)
	val2, ok2 := myAtoi64(val2Str)

	if !ok1 || !ok2 {
		return
	}

	var result int64
	overflow := false

	switch op {
	case "+":
		if (val2 > 0 && val1 > MaxInt64-val2) || (val2 < 0 && val1 < MinInt64-val2) {
			overflow = true
		} else {
			result = val1 + val2
		}

	case "-":
		v2Neg := -val2
		if (v2Neg > 0 && val1 > MaxInt64-v2Neg) || (v2Neg < 0 && val1 < MinInt64-v2Neg) {
			overflow = true
		} else {
			result = val1 - val2
		}

	case "*":
		if val1 != 0 && val2 != 0 {
			result = val1 * val2
			if val1 == MinInt64 && val2 == -1 {
				overflow = true
			} else if result/val1 != val2 {
				overflow = true
			}
		} else {
			result = 0
		}

	case "/":
		if val2 == 0 {
			printStr("No division by 0")
			return
		}
		if val1 == MinInt64 && val2 == -1 {
			overflow = true
		} else {
			result = val1 / val2
		}

	case "%":
		if val2 == 0 {
			printStr("No modulo by 0")
			return
		}
		result = val1 % val2

	default:
		return
	}

	if overflow {
		return
	}

	os.Stdout.Write(itoa(result))
	os.Stdout.Write([]byte{'\n'})
}
