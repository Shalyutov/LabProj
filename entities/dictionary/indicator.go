package dictionary

type IntegerIndicator struct {
	Name    string
	Measure string
	High    float64
	Low     float64
}

type BinaryIndicator struct {
	Name      string
	Reference bool
}

type StringIndicator struct {
	Name      string
	Reference string
}
