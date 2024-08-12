package secao11generics

import "fmt"

func Main() {
	// Without generics
	slice := []int{5, 1, 2, 3}
	newSlice := reverseInt(slice)
	fmt.Println(newSlice)

	slice2 := []string{"a", "teste", "josé"}
	newSlice2 := reverseString(slice2)
	fmt.Println(newSlice2)

	// With generics
	slice3 := []int{6, 24, 45, 10, 3}
	newSlice3 := reverseGenerics(slice3)
	fmt.Println(newSlice3)

	slice4 := []string{"teste2", "ola", "yoda"}
	newSlice4 := reverseGenerics(slice4)
	fmt.Println(newSlice4)
}

func reverseInt(slice []int) []int {
	newInts := make([]int, len(slice))

	lastUnallocatedPosition := len(newInts) - 1
	for i := 0; i < len(slice); i++ {
		newInts[lastUnallocatedPosition] = slice[i]
		lastUnallocatedPosition--
	}

	return newInts
}

func reverseString(slice []string) []string {
	newStrings := make([]string, len(slice))

	lastUnallocatedPosition := len(newStrings) - 1
	for i := 0; i < len(slice); i++ {
		newStrings[lastUnallocatedPosition] = slice[i]
		lastUnallocatedPosition--
	}

	return newStrings
}

func reverseGenerics[T int | string](slice []T) []T {
	newSlice := make([]T, len(slice))

	lastUnallocatedPosition := len(newSlice) - 1
	for i := 0; i < len(slice); i++ {
		newSlice[lastUnallocatedPosition] = slice[i]
		lastUnallocatedPosition--
	}

	return newSlice
}

type constraintCustom interface {
	int | string
}

func reverseConstraint[T constraintCustom](slice []T) []T {
	newSlice := make([]T, len(slice))

	lastUnallocatedPosition := len(newSlice) - 1
	for i := 0; i < len(slice); i++ {
		newSlice[lastUnallocatedPosition] = slice[i]
		lastUnallocatedPosition--
	}

	return newSlice
}
