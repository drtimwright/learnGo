package tip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanary(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestZeroTip(t *testing.T) {
	total, tipAmount := CalculateTip(100, 0)
	assert.Equal(t, tipAmount, 0)
	assert.Equal(t, total, 100)
}

func TestOneHundredPercentTip(t *testing.T) {
	total, tipAmount := CalculateTip(100, 100)
	assert.Equal(t, tipAmount, 100)
	assert.Equal(t, total, 200)
}

func TestFifteenPercentTip(t *testing.T) {
	total, tipAmount := CalculateTip(100, 15)
	assert.Equal(t, tipAmount, 15)
	assert.Equal(t, total, 115)
}

func TestTipRounding(t *testing.T) {
	total, tipAmount := CalculateTip(10, 15)
	assert.Equal(t, tipAmount, 2)
	assert.Equal(t, total, 12)
}

func TestTipRoundingUp(t *testing.T) {
	total, tipAmount := CalculateTip(1, 15)
	assert.Equal(t, tipAmount, 1)
	assert.Equal(t, total, 2)
}
