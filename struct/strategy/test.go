package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	op := &Operator{}

	op.setStrategy(&add{})
	result := op.calculate(1, 2)
	fmt.Println("add:", result)

	op.setStrategy(&reduce{})
	result = op.calculate(2, 1)
	fmt.Println("reduce:", result)
}
