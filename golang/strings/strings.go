package strings

import (
	"fmt"
	"strings"
)

// 文本字段
func Field() {
	ss := strings.Fields("  foo bar  baz   ")
	fmt.Println("[]slice =", ss)
}

func Contains() {
	str := "abcdfeg"
	cons1 := strings.Contains(str, "a")
	cons2 := strings.Contains(str, "h")
	fmt.Println("a? ", cons1)
	fmt.Println("b? ", cons2)

}

func Joins() {
	s := []string{"foo", "bar", "baz"}
	prt := strings.Join(s, ", ")
	fmt.Println(prt)
}

func Count() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
}
