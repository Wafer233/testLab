package channel_select

import (
	"fmt"
	"testing"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

// select可以同时监听一个或多个channel，直到其中一个channel ready
func TestSelect(t *testing.T) {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go test1(output1)
	go test2(output2)
	// 用select监控
	select {
	case s1 := <-output1:
		fmt.Println("s1 = ", s1)
	case s2 := <-output2:
		fmt.Println("s2 = ", s2)
	}
}

// 如果多个channel同时ready，则随机选择一个执行
func TestSelect2(t *testing.T) {
	// 创建2个管道
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	go func() {

		intChan <- 1
	}()
	go func() {
		//time.Sleep(2 * time.Second)
		stringChan <- "hello"
	}()
	select {
	case value := <-intChan:
		fmt.Println("int:", value)
	case value := <-stringChan:
		fmt.Println("string:", value)
	}
	fmt.Println("main结束")
}

func TestSelect3(t *testing.T) {
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

// 可以用于判断管道是否存满
func TestFull(t *testing.T) {
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	//取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
