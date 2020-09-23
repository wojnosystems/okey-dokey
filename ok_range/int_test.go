package ok_range

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInt(t *testing.T) {
	cases := map[string]struct {
		start, end   int
		eStart, eEnd int
	}{
		"in-order": {
			start:  4,
			end:    6,
			eStart: 4,
			eEnd:   6,
		},
		"reversed": {
			start:  6,
			end:    4,
			eStart: 4,
			eEnd:   6,
		},
		"equal": {
			start:  4,
			end:    4,
			eStart: 4,
			eEnd:   4,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			actual := IntBetween(c.start, c.end)
			assert.Equal(t, c.eStart, actual.Start())
			assert.Equal(t, c.eEnd, actual.End())
		})
	}
}
