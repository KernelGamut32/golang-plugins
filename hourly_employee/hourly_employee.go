package main

type hourlyEmployee struct{}

func (h *hourlyEmployee) ProcessPay(pay_rate float64) float64 {
	return 40 * pay_rate
}

var Employee hourlyEmployee
