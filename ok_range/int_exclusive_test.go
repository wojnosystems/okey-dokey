package ok_range

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIntExclusive(t *testing.T) {
	cases := map[string]struct {
		start, end int
		expected   error
	}{
		"not equal": {
			start: 3,
			end:   4,
		},
		"equal": {
			start:    4,
			end:      4,
			expected: ErrExclusiveRangeWasEmpty,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			_, err := IntBetweenExclusive(c.start, c.end)
			if c.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, c.expected.Error())
			}
		})
	}
}
