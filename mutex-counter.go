package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	v map[string]int
	sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.v[key]
}

func main() {
	v := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go v.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(v.Value("somekey"))
}
