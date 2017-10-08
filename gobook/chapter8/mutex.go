package main

import (
	"fmt"
	"sync"
	"time"
)

// MapCounter is a thread safe data structure with mutex inside
type MapCounter struct {
	v      map[string]int
	mutext sync.Mutex
}

// Inc locking for read/write
func (c *MapCounter) Inc(key string) {
	c.mutext.Lock()
	c.v[key]++
	c.mutext.Unlock()
}

func main() {
	counter := MapCounter{v: make(map[string]int)}

	for i := 0; i < 10000; i++ {
		go counter.Inc("somekey")
	}

	time.Sleep(1 * time.Second)
	fmt.Println(counter.v["somekey"])
}
