package strategy

// 策略模式（Strategy Pattern）定义一组算法，将每个算法都封装起来，并且使它们之间可以互换。

type IStrategy interface {
	do(int, int) int
}

type add struct {
}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

func (p *Operator) setStrategy(strategy IStrategy) {
	p.strategy = strategy
}

func (p *Operator) calculate(a, b int) int {
	return p.strategy.do(a, b)
}
