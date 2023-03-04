package safe

import "sync"

type Singleton struct {
}

var ins *Singleton
var mu sync.Mutex

func GetInsOr() *Singleton {
	if ins == nil {
		mu.Lock()
		if ins == nil {
			ins = &Singleton{}
		}
		mu.Unlock()
	}
	return ins
}
