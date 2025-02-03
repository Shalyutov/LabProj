package dictionary

type IntegerIndicator struct {
	Name    string
	Comment string
	Measure Measure
	High    float64
	Low     float64
}

type BinaryIndicator struct {
	Name      string
	Comment   string
	Reference bool
}

type StringIndicator struct {
	Name      string
	Comment   string
	Reference string
}
