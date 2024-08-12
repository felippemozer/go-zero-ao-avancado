package secao14goroutines

import (
	"fmt"
	"sync"
	"time"
)

func callDB(wg *sync.WaitGroup) {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Finalizado callDB")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Finalizado callAPI")
	wg.Done()
}

func internalProcess(wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Finalizado internalProcess")
	wg.Done()
}
