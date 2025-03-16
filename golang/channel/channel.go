package channel

import "fmt"

//对一个关闭的通道再发送值就会导致panic。
//对一个关闭的通道进行接收会一直获取值直到通道为空。
//对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//关闭一个已经关闭的通道会导致panic。

func Recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
