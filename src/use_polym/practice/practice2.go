package practice


type Advertisement struct {
	AdvName string
	CPC int
	NoOfClicks int
}

func (e Advertisement) calculate() int {
	return e.CPC * e.NoOfClicks
}

func (e Advertisement) source() string {
	return e.AdvName
}