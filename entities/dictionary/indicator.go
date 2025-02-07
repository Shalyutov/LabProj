package dictionary

type IntegerIndicator struct {
	Name    string
	Measure string
	High    float64
	Low     float64
	Id      int
}

type BinaryIndicator struct {
	Name      string
	Reference bool
	Id        int
}

type StringIndicator struct {
	Name      string
	Reference string
	Id        int
}
