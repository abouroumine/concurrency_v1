package main

import (
	"fmt"
	"sync"
)

func increaseByOneWithoutMutex(x *int, wg *sync.WaitGroup) {
	*x = *x + 1
	wg.Done()
}

func increaseByOneWithMutex(x *int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*x = *x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	xWith := 0
	xWithout := 0
	for i := 0; i < 50000; i++ {
		w.Add(2)
		go increaseByOneWithoutMutex(&xWithout, &w)
		go increaseByOneWithMutex(&xWith, &w, &m)
	}
	w.Wait()
	fmt.Println("Value without Mutex = ", xWithout)
	fmt.Println("Value With Mutex = ", xWith)
}
