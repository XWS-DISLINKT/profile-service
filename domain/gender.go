package domain

type Gender int8

const (
	Male Gender = iota
	Female
	Other
)

func (gender Gender) String() string {
	switch gender {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case Other:
		return "Other"
	}
	return "Unknown"
}
