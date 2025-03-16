package channel

import (
	"fmt"
	"testing"
)

// err！ 无缓冲必须又接受
func TestChannel(t *testing.T) {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}

func TestChannelv1(t *testing.T) {
	ch := make(chan int)
	go Recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

func TestChannelv2(t *testing.T) {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
}

func TestClose(t *testing.T) {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}
