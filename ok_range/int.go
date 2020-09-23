package ok_range

// Int provides a range such that Start() is always less than or equal to End()
type Int struct {
	startAt int
	endAt   int
}

func IntBetween(startAt int, endAt int) Int {
	if startAt <= endAt {
		return Int{
			startAt: startAt,
			endAt:   endAt,
		}
	}
	return Int{
		startAt: endAt,
		endAt:   startAt,
	}
}

func (i Int) Start() int {
	return i.startAt
}

func (i Int) End() int {
	return i.endAt
}
