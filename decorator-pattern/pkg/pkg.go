package pkg

// Assume this as an external package
type FeeProvider interface {
	ProcessingFee() float64
	TransferFee() float64
}

type Transfer struct {
	Amount float64
}

func (c *Transfer) ProcessingFee() float64 {
	return 50
}

func (c *Transfer) TransferFee() float64 {
	return c.Amount * 0.01
}
