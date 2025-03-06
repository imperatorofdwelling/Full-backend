package sort

type Sort int // @name Sort

const (
	Nil Sort = iota
	Old
	New
	HighlyRecommended
	LowlyRecommended
)

// String method to convert Sort to its string representation
func (s Sort) String() string {
	return [...]string{
		"Nil",
		"Old",
		"New",
		"Highly Recommended",
		"Lowly Recommended",
	}[s]
}
