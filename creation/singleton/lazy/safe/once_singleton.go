package safe

import "sync"

type OnceSingleton struct {
}

var onceIns *OnceSingleton
var once sync.Once

func GetOnceInsOr() *OnceSingleton {
	once.Do(func() {
		onceIns = &OnceSingleton{}
	})
	return onceIns
}
