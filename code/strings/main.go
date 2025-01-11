package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "apple,banana,grapes,mango"
	parts := strings.Split(str, ",")
	fmt.Println(parts)

	fmt.Println(strings.Count(str, "a"))

	str2 := "    Hello, World!    "
	fmt.Println(strings.TrimSpace(str2))

	str3 := "Subham"
	str4 := "Mohanty"
	res := strings.Join([]string{str3,"Hello",str4},"----")
	fmt.Println(res)
}
