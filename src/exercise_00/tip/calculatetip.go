package tip

import "math"

func CalculateTip(amountInCents int64, tipPercentage int64) (int64, int64) {
	tipAmount := int64(math.Ceil(float64(amountInCents) * float64(tipPercentage) / 100.0))
	return amountInCents + tipAmount, tipAmount
}
