package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //用于等待一组协程完成
var counter int

func main() {
	const grs = 2

	var mu sync.Mutex // 互斥锁保护计数器

	for i := 0; i < grs; i++ {
		wg.Add(1)
		go func() {
			mu.Lock() // 在修改变量之前锁定
			{
				counter++
			}
			mu.Unlock() //修改后解锁
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}
