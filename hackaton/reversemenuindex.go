package piscine

func ReverseMenuIndex(menu []string) []string {
	length := 0
	for range menu {
		length++
	}

	reversedMenu := make([]string, length)

	for i := 0; i < length; i++ {
		reversedMenu[i] = menu[length-1-i]
	}

	return reversedMenu
}
