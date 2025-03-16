package sort

import (
	"fmt"
	"slices"
	"sort"
)

// go 1.22 update
func SliceSort() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

func Sorts() {
	s1 := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s1)
	fmt.Println(s1)

	s3 := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s3)
	fmt.Println(s3)

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)

	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	//排序条件和排序对象不一定是同一个东西，非常自由的方法
	//e.g 我想排序 []num,使用 map[int]int 中的 value也是可以的，这样我就能给map排序（根据map的value值）
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}
