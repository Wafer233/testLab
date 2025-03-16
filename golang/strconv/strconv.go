package strconv

import (
	"fmt"
	"strconv"
)

func Atoi() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
}

//【扩展阅读】这是C语言遗留下的典故。C语言中没有string类型而是用字符数组(array)表示字符串，所以Itoa对很多C系的程序员很好理解。

func Itoa() {
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
}

func Parse() {
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-2", 10, 64)
	u, _ := strconv.ParseUint("2", 10, 64)
	fmt.Println("b:", b)
	fmt.Println("f:", f)
	fmt.Println("i:", i)
	fmt.Println("u:", u)
}

func Format() {
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Println("bool", s1)
	fmt.Println("float64", s2)
	fmt.Println("int64", s3)
	fmt.Println("unit64", s4)
}
