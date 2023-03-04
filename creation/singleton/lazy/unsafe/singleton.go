package unsafe

type Singleton struct {
}

var ins *Singleton

func GetInsOr() *Singleton {
	if ins == nil {
		ins = &Singleton{}
	}
	return ins
}
