package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	var sum = 0
	for {
		var s string
		fmt.Scanln(&s)
		if s == "" {
			break
		}

		var num string

		for i := 0; i < len(s); i++ {
			if unicode.IsLetter(rune(s[i])) {

			} else {
				num = num + string(s[i])
				break
			}
		}

		for i := len(s) - 1; i >= 0; i-- {
			if unicode.IsLetter(rune(s[i])) {

			} else {
				num = num + string(s[i])
				break
			}
		}

		i, err := strconv.Atoi(num)
		if err != nil {
			// ... handle error
			panic(err)
		}
		sum = sum + i

	}
	fmt.Println(sum)
}
