package sort

// Sort represents the type for sorting options
type Sort string // @name Sort

// Enumerate the available sorting options
const (
	Nil               Sort = "Nil"
	Old               Sort = "Old"
	New               Sort = "New"
	HighlyRecommended Sort = "Highly Recommended"
	LowlyRecommended  Sort = "Lowly Recommended"
)

// String method to convert Sort to its string representation
func (s Sort) String() string {
	return string(s)
}

// AllSorts returns a slice of all available sorting options
func AllSorts() []Sort {
	return []Sort{
		Nil,
		Old,
		New,
		HighlyRecommended,
		LowlyRecommended,
	}
}
