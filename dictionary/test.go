package dictionary

type Test struct {
	Name        string
	Aliases     []string
	Indicators  []Indicator
	Services    []Service
	Cases       map[Biomaterial]Supply
	IsSeparated bool
	Price       float64
}
