package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/jrauljperez07/go-routine/data"
)

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readBook(i, wg, m)
	}

	wg.Wait()
	duration := time.Since(start).Milliseconds()
	fmt.Printf("Total time: %dms\n", duration)
}

func showGoRoutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutine #%d with delay %dms\n", id, delay)

	time.Sleep(time.Duration(delay) * time.Millisecond)

	wg.Done()
}

func readBook(id int, wg *sync.WaitGroup, m *sync.Mutex) {
	data.FinishBook(id, m)

	delay := rand.Intn(500)
	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}
