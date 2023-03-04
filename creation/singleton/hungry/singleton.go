package hungry

type Singleton struct {
}

var ins = &Singleton{}

func GetInsOr() *Singleton {
	return ins
}
