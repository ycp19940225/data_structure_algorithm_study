package main

import "fmt"

func main() {
	testArray()
}

func testArray() {
	var array1 [10]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	for i := 0; i < len(array1); i++ {
		fmt.Printf("%v\n", array1[i])
	}

	arrayStr := [5]string{"s", "s2", "s3"}

	for i := 0; i < len(arrayStr); i++ {
		fmt.Printf("%v\n", arrayStr[i])
	}

	arrayMap := new([10]map[string]string)
	arrayMap[0] = map[string]string{1, 2}
}
