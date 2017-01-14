package goutils

import (
	"fmt"
	"os"
	"strconv"
)

/*
CompareTwoArraysOfStrings is an implementation to campare two elements of type Array.
*/
func CompareTwoArraysOfStrings(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}

	x := 0
	evaluation := true

	for x < len(first) && evaluation == true {
		if first[x] != second[x] {
			evaluation = false
		}
		x = x + 1
	}

	return evaluation
}

/*CompareTwoArraysOfInt is an implementation that return true if the
elements of an array of int are equal*/
func CompareTwoArraysOfInt(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}

	x := 0
	evaluation := true

	for x < len(first) && evaluation == true {
		if first[x] != second[x] {
			evaluation = false
		}
		x = x + 1
	}

	return evaluation
}

/*SumOfAnArrayIntElements should sum the elements of an array
with recursive behaviour.*/
func SumOfAnArrayIntElements(list []int) []int {
	x := 0
	var result []int
	//result := make([]int, width, width)
	suma := 0
	for x < len(list) {
		suma = suma + list[x]
		result = append(result, suma)
		x = x + 1
	}
	return result
}

/*SumArrayOfStringsAsInt sum an array of strings as int on a recursive way*/
func SumArrayOfStringsAsInt(list []string) []string {
	x := 0
	var result []string
	//result := make([]int, width, width)
	suma := 0
	for x < len(list) {
		convertedInt := ParseStringToInt(list[x])
		suma = suma + convertedInt
		result = append(result, strconv.Itoa(suma))
		x = x + 1
	}
	return result
}

/*ParseStringToInt A string to int parser*/
func ParseStringToInt(num string) int {
	result, err := strconv.Atoi(num)

	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return result
}

/*ParseIntToString A string to int parser*/
func ParseIntToString(num int) string {
	result := strconv.Itoa(num)

	return result
}
