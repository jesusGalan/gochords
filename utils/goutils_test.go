package goutils_test

import (
	"testing"

	"github.com/jesusGalan/goutils"
)

func TestCompareArraysOfStringsError(t *testing.T) {
	firstArray := []string{"hola", "mundo"}
	secondArray := []string{"kata", "compare"}
	result := goutils.CompareTwoArraysOfStrings(firstArray, secondArray)

	//if result == true {
	if result {
		t.Error("Yours Arrays are equal.")
	}

}

func TestCompareArraysOfStringPass(t *testing.T) {
	firstArray := []string{"hola", "mundo"}
	secondArray := []string{"hola", "mundo"}
	result := goutils.CompareTwoArraysOfStrings(firstArray, secondArray)

	if !result {
		t.Error("Your Arrays are not equal.")
	}
}

func TestCompareWidthOfArrays(t *testing.T) {
	firstArray := []string{"hola", "mundo"}
	secondArray := []string{"hola", "mundo", "kata"}
	result := goutils.CompareTwoArraysOfStrings(firstArray, secondArray)

	if result {
		t.Error("Your array have the same width.")
	}
}

func TestCompareArraySecondMember(t *testing.T) {
	firstArray := []string{"hola", "mundo"}
	secondArray := []string{"hello", "mundo"}
	result := goutils.CompareTwoArraysOfStrings(firstArray, secondArray)

	if result {
		t.Error("Your arrays are not equal.")
	}

}

func TestCompareArraysOfIntError(t *testing.T) {
	firstArray := []int{1, 2}
	secondArray := []int{1, 2, 3}
	result := goutils.CompareTwoArraysOfInt(firstArray, secondArray)

	//if result == true {
	if result {
		t.Error("Yours Arrays are equal.")
	}

}

func TestCompareArraysOfIntMembersPass(t *testing.T) {
	firstArray := []int{1, 2, 3}
	secondArray := []int{1, 2, 3}
	result := goutils.CompareTwoArraysOfInt(firstArray, secondArray)

	//if result == true {
	if !result {
		t.Error("Yours Arrays elements are not equal.")
	}

}

func TestCompareWidthOfArraysOfInt(t *testing.T) {
	firstArray := []int{1, 2}
	secondArray := []int{1, 7, 8}
	result := goutils.CompareTwoArraysOfInt(firstArray, secondArray)

	//if result == true {
	if result {
		t.Error("Yours Arrays have not the same width")
	}

}

func TestSumOfElementsOfAnArrayOfIntsError(t *testing.T) {
	array := []int{1, 2, 2}
	expected := []int{1, 3, 6}
	result := goutils.SumOfAnArrayIntElements(array)
	flag := goutils.CompareTwoArraysOfInt(result, expected)

	if flag {
		t.Error("The proccess have not errors.")
	}
}

func TestSumOfElementsOfAnArrayOfIntsPass(t *testing.T) {
	array := []int{2, 2, 2}
	expected := []int{2, 4, 6}
	result := goutils.SumOfAnArrayIntElements(array)
	flag := goutils.CompareTwoArraysOfInt(result, expected)

	if !flag {
		t.Error("The proccess have errors.")
	}
}
