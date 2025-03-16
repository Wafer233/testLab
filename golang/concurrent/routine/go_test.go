package routine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 当开启一个main时，可以理解为开启了一个go routine，这时候并行另外一个go routine，当main结束 全部结束
func TestConcurrent(t *testing.T) {
	go Hello()

	fmt.Println("main goroutine done!")
}

// 因此给sleep来等待go hello结束
func TestConcurrentV1(t *testing.T) {
	// 启动一个goroutine去执行hello函数
	go Hello()

	fmt.Println("test goroutine done!")
	time.Sleep(time.Second)
}

// go rountine 随机
func TestHelloV2(t *testing.T) {

	for i := 0; i < 10; i++ {
		go Count(i)
	}
	fmt.Println("test goroutine done!")
	//time.Sleep(time.Second * 5)
}

func TestRuntime(t *testing.T) {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func TestRuntimeV2(t *testing.T) {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}
