package employee

type Employee interface {
	ProcessPay(pay_rate float64) float64
}
