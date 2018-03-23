package main

import (
	"fmt"
	"time"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	return &Mutex{make(chan struct{}, 1)}
}

func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

func (m *Mutex) Unlock() {
	<-m.ch
}

func (m *Mutex) TryLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case m.ch <- struct{}{}:
		timer.Stop()
		return true
	case <-time.After(timeout):
	}
	return false
}

func (m *Mutex) IsLocked() bool {
	return len(m.ch) > 0
}

func main() {
	m := NewMutex()
	ok := m.TryLock(time.Second)
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock(time.Second)
	fmt.Printf("locked %v\n", ok)
}
