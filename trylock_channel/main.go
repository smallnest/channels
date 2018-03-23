package main

import (
	"fmt"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	// using n i like Waitgroup
	return &Mutex{make(chan struct{}, 1)}
}

func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

func (m *Mutex) Unlock() {
	<-m.ch
}

func (m *Mutex) TryLock() bool {
	select {
	case m.ch <- struct{}{}:
		return true
	default:
	}
	return false
}

func (m *Mutex) IsLocked() bool {
	return len(m.ch) > 0
}

func main() {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Printf("locked v %v\n", ok)
	ok = m.TryLock()
	fmt.Printf("locked %v\n", ok)
}
