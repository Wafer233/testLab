package slices

import (
	"fmt"
	"slices"
)

func Contain() {
	numbers := []int{0, 1, 2, 3}
	fmt.Println(slices.Contains(numbers, 2))
	fmt.Println(slices.Contains(numbers, 4))

}

func Sort() {
	smallInts := []int8{0, 42, -10, 8}
	slices.Sort(smallInts)
	fmt.Println(smallInts)
}
