package main

type salariedEmployee struct{}

func (h *salariedEmployee) ProcessPay(pay_rate float64) float64 {
	return (pay_rate * 2080) / 52
}

var Employee salariedEmployee
