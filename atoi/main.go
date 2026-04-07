package main

import (
	"fmt"
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))      // 125
	fmt.Println(piscine.AtoiBase("1111101", "01"))          // 125
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF")) // 125
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))          // 125
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))         // 0
}
