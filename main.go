package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := lenAndUpper("hoho")
	fmt.Println(totalLength, upperName)
	total := superAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)
	fmt.Println(canIDrink(20))

}

func superAdd(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

	// switch {
	// case age < 18:
	// 	return false
	// case age == 18:
	// 	return true
	// }
	// return false
}
