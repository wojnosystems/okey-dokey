package ok_range

import (
	"errors"
	"log"
)

// IntExclusive provides a range such that Start() is always less than or equal to End()
// Further more, it is a run-time error if you try to use the same value for start and end.
type IntExclusive struct {
	Int
}

var ErrExclusiveRangeWasEmpty = errors.New("startAt was equal to endAt in an exclusive range")

func IntBetweenExclusive(startAt int, endAt int) (i IntExclusive, err error) {
	i = IntExclusive{
		Int: IntBetween(startAt, endAt),
	}
	if startAt == endAt {
		err = ErrExclusiveRangeWasEmpty
	}
	return
}

// MustIntBetweenExclusive is just like IntBetweenExclusive, but never returns an error.
// If startAt == endAt, it will panic instead
func MustIntBetweenExclusive(startAt int, endAt int) (i IntExclusive) {
	var err error
	i, err = IntBetweenExclusive(startAt, endAt)
	if err != nil {
		log.Panic(err)
	}
	return
}
