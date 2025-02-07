package dictionary

type Test struct {
	Name              string
	Aliases           []string
	IntegerIndicators []IntegerIndicator
	BinaryIndicators  []BinaryIndicator
	StringIndicators  []StringIndicator
	Services          []Service
	Cases             []Supply
	IsSeparated       bool
	Price             float64
	Id                int
}
