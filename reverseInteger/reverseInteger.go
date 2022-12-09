package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func reverse(x int) int {
	fmt.Println("var1 = ", reflect.TypeOf(x))

	slice := transformNumberOnSlice(x)
	var str string
	for _, v := range slice {
		str += strconv.Itoa(v)
	}
	result, err := strconv.Atoi(str)

	if err != nil {
	}

	if result < -int(math.Pow(2, 31)) {
		return 0
	}

	if result > int(math.Pow(2, 31))-1 {
		return 0
	}

	return result

}

func transformNumberOnSlice(n int) []int {
	var aux []int
	a := 0
	acceptZero := false
	for n != 0 {
		moduleBy10 := n % 10
		if moduleBy10 < 0 && a > 0 {
			moduleBy10 = n % 10 * -1
		}
		if moduleBy10 != 0 {
			acceptZero = true
			a++
		}
		if acceptZero {
			aux = append(aux, moduleBy10)
		}

		n /= 10
	}
	return aux
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(901000))
	fmt.Println(reverse(1534236469))
}
