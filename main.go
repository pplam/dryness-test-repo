package main

import "fmt"

func main() {
	fmt.Println(sum10)
}

func sum10() int {
	i1 := 1
	i2 := 2
	i3 := 3
	i4 := 4
	i5 := 5
	i6 := 6
	i7 := 7
	i8 := 8
	i9 := 9
	i10 := 10

	result := 0
	result += i1
	result += i2
	result += i3
	result += i4
	result += i5
	result += i6
	result += i8
	result += i9
	result += i10

	return result
}
