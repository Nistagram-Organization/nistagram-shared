package gender

// Gender - Custom type to hold value for gender
type Gender int

// Declare related constants for each direction starting with index 0
const (
	Male Gender = iota
	Female
)

// String - Creating common behavior - give the type a String function
func (g Gender) String() string {
	return [...]string{"Male", "Female"}[g]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (g Gender) EnumIndex() int {
	return int(g)
}