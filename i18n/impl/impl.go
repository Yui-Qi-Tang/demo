package impl

import (
	"fmt"
	"sync"

	"iprotector.github.com/i18n"
)

type t struct {
}

func (t) Translate() {
	fmt.Println("i am t")
}

func init() {
	i18n.I = t{}
}

type T1 struct {
	mu sync.RWMutex
}

func (t *T1) Set() {
	t.mu.Lock()
	defer t.mu.Unlock()
}
