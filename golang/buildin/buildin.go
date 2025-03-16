package buildin

import "fmt"

func Append() {
	//slice = append(slice, elem1, elem2)
	//slice = append(slice, anotherSlice)
	s1 := []int{1, 2, 3}
	i := append(s1, 4)
	fmt.Println(i)

	s2 := []int{7, 8, 9}
	i2 := append(s1, s2...)
	fmt.Println(i2)
}

func Len() {
	//数组、切片、字符串、通道的长度

	s1 := "hello world"
	i := len(s1)
	fmt.Println(i)

	s2 := []int{1, 2, 3}
	fmt.Println(s2)
}

func Print() {
	name := "wafer"
	age := 20
	print(name, " ", age, "\n")
	fmt.Println("---------")
	println(name, " ", age)
}

func Panic() {
	defer fmt.Println("after panic ...")
	panic("panic ...")
	fmt.Println("end...")
}

func Recover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...")
		}
	}()
	panic("panic...")
}

/*
new和make区别：

make只能用来分配及初始化类型为slice，map，chan的数据；
new可以分配任意类型的数据；
new分配返回的是指针，即类型*T；make返回引用，即T；
new分配的空间被清零，make分配后，会进行初始化
*/

func New() {
	b := new(bool)
	fmt.Println(*b)
	i := new(int)
	fmt.Println(*i)
	s := new(string)
	fmt.Println("strings:", *s)
}

/*
例如：make([]int, 10 , 100)

说明：分配一个有100个int的数组，然后创建一个长度为10，容量为100的slice结构，该slice引用包含前10个元素的数组。
对应的，new([]int)返回一个指向新分配的，被置零的slice结构体的指针，即指向值为nil的slice的指针。

内建函数make(T, args)与new(T)的用途不一样。它只用来创建slice、map和channel，
并且返回一个初始化的（而不是置零），类型为T的值（而不是*T）。之所以有所不同，是因为这三个类型的背后引用了使用前必须初始化的数据结构。
例如，slice是一个三元描述符，包含一个指向数据（在数组中）的指针，长度，以及容量，在这些项被初始化之前，slice都是nil的。
对于slice，map和channel，make初始化这些内部数据结构，并准备好可用的值。
*/

func Make() {
	var p *[]int = new([]int)     // allocates slice structure; *p == ni; rarely useful
	var v []int = make([]int, 10) // the slice v now refers to a new array of 100 ints

	fmt.Println(p)
	fmt.Println(v)

	var p1 *[]int = new([]int)
	*p1 = make([]int, 5, 10)

	// Idiomatic: 习惯的做法
	v1 := make([]int, 10)

	fmt.Println(p1)
	fmt.Println(v1)
}
