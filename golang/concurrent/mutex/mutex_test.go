package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 没有锁
func addWithoutMutex() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

func TestMutex(t *testing.T) {
	wg.Add(2)
	go addWithoutMutex()
	go addWithoutMutex()
	wg.Wait()
	fmt.Println(x)
}

//6.1 互斥锁

func addWithMutex() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func TestMutex2(t *testing.T) {
	wg.Add(2)
	go addWithMutex()
	go addWithMutex()
	wg.Wait()
	fmt.Println(x)
}

// 6.2 读写互斥锁互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，
//当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择
//读写锁在Go语言中使用sync包中的RWMutex类型。

//读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，
//如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	//lock.Lock() // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	//lock.Unlock() // 解互斥锁
	wg.Done()
}

func read() {
	//lock.Lock() // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	//lock.Unlock() // 解互斥锁
	wg.Done()
}

// 1.86S -> 107 ms
func TestRWMutex(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
