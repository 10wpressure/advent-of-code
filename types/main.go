package types

type Range struct {
	Start int64
	End   int64
}

func (r *Range) Contains(val int64) bool {
	return val >= r.Start && val < r.End
}

func (r *Range) ContainsNotIncluding(val int64) bool {
	return val > r.Start && val < r.End
}
