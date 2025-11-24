package dictionary

type Indicator struct {
	Name            string
	Measure         string
	HighReference   float64
	LowReference    float64
	BinaryReference bool
	StringReference string
	Id              int
}
