package dictionary

type Test struct {
	Name        string
	Aliases     []string
	Indicators  []Indicator
	Services    []Service
	Cases       []Supply
	IsSeparated bool
	Price       float64
}
