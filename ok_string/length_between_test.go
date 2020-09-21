package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"okey-dokey/bad"
	"testing"
)

func TestLengthBetween_Validate(t *testing.T) {
	cases := map[string]struct {
		atLeast  int
		atMost   int
		input    *string
		expected string
	}{
		"ok": {
			atLeast:  4,
			atMost:   6,
			input:    addrOf("exact"),
			expected: "",
		},
		"too long": {
			atLeast:  4,
			atMost:   6,
			input:    addrOf("too long"),
			expected: "must have at least 4 and at most 6 characters, but had 8",
		},
		"too short": {
			atLeast:  4,
			atMost:   6,
			input:    addrOf("sht"),
			expected: "must have at least 4 and at most 6 characters, but had 3",
		},
		"nil": {
			atLeast: 4,
			atMost:  6,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := LengthBetween{
				AtLeast: c.atLeast,
				AtMost:  c.atMost,
			}
			validationErrors := bad.SliceMessageReceiver{}
			validator.Validate(c.input, &validationErrors)
			if len(c.expected) != 0 {
				require.Len(t, validationErrors, 1)
				assert.Equal(t, c.expected, validationErrors[0])
			} else {
				assert.Equal(t, 0, len(validationErrors))
			}
		})
	}
}
