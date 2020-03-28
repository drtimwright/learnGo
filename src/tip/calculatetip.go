package tip

import "math"

func CalculateTip(amountInCents int, tipPercentage int) (int, int) {
	tipAmount := int(math.Round(float64(amountInCents) * float64(tipPercentage) / 100.0))
	return amountInCents + tipAmount, tipAmount
}

