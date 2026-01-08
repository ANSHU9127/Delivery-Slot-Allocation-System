package main

import (
    "fmt"
    "sync"
)

type SlotManager struct {
    slots int
    mu    sync.Mutex
}

func (s *SlotManager) Allocate(wg *sync.WaitGroup, id int) {
    defer wg.Done()
    s.mu.Lock()
    defer s.mu.Unlock()

    if s.slots > 0 {
        s.slots--
        fmt.Println("Slot allocated to request", id)
    } else {
        fmt.Println("No slots left for request", id)
    }
}

func main() {
    manager := SlotManager{slots: 5}
    var wg sync.WaitGroup

    for i := 1; i <= 10; i++ {
        wg.Add(1)
        go manager.Allocate(&wg, i)
    }

    wg.Wait()
}