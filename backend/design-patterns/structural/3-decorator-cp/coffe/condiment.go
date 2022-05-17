package coffe

type Mocha struct {
	// TODO: answer here
	Coffe Coffe
}

func (m Mocha) GetCost() float64 {
	return m.Coffe.GetCost() + 1.00 // TODO: replace this
}

func (m Mocha) GetDescription() string {
	return m.Coffe.GetDescription() + ", Mocha" // TODO: replace this
}

type Whipcream struct {
	// TODO: answer here
	Coffe Coffe
}

func (w Whipcream) GetCost() float64 {
	return w.Coffe.GetCost() + 0.10 // TODO: replace this
}

func (w Whipcream) GetDescription() string {
	return w.Coffe.GetDescription() + ", Whipcream" // TODO: replace this
}
