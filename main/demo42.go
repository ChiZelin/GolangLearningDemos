package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex // 读写锁
	count int          // 共享计数器
}

// Read 允许多个 goroutine 并发读
func (c *Counter) Read() int {
	c.mu.RLock()         // 获取读锁
	defer c.mu.RUnlock() // 释放读锁
	return c.count
}

// Write 同一时间只允许一个 goroutine 写
func (c *Counter) Write() {
	c.mu.Lock()         // 获取写锁
	defer c.mu.Unlock() // 释放写锁
	c.count++           // 修改共享资源
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// 启动 10 个读 goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
				fmt.Printf("Reader %d: count = %d\n", id, counter.Read())
			}
		}(i)
	}

	// 启动 2 个写 goroutine
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				counter.Write()
				fmt.Printf("Writer %d: incremented count to %d\n", id, counter.Read())
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Final count:", counter.Read())
}
