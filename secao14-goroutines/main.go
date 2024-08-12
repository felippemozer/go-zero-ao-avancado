package secao14goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Main() {
	// wait groups

	var wg sync.WaitGroup
	wg.Add(3)
	go callDB(&wg)
	go callAPI(&wg)
	go internalProcess(&wg)

	wg.Wait()

	// mutex
	var m sync.Mutex
	x := 0
	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			x++
			m.Unlock()
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(x)

	// channels
	ch := make(chan int, 100)
	go setList(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func setList(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
