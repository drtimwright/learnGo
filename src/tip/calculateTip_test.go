package tip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

import "math"

func ComputeTip(amountInCents int, tipPercentage int) (int, int) {
	tipAmount := int(math.Round(float64(amountInCents) * float64(tipPercentage) / 100.0))
	return amountInCents + tipAmount, tipAmount
}

func TestZeroTip(t *testing.T) {
	total, tipAmount := ComputeTip(100, 0)
	assert.Equal(t, tipAmount, 0)
	assert.Equal(t, total, 100)
}

func TestOneHundredPercentTip(t *testing.T) {
	total, tipAmount := ComputeTip(100, 100)
	assert.Equal(t, tipAmount, 100)
	assert.Equal(t, total, 200)
}

func TestFifteenPercentTip(t *testing.T) {
	total, tipAmount := ComputeTip(100, 15)
	assert.Equal(t, tipAmount, 15)
	assert.Equal(t, total, 115)
}

func TestTipRounding(t *testing.T) {
	total, tipAmount := ComputeTip(10, 15)
	assert.Equal(t, tipAmount, 2)
	assert.Equal(t, total, 12)
}
