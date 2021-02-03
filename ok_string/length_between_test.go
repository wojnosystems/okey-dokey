package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_range"
	"testing"
)

func TestLengthBetween_Validate(t *testing.T) {
	cases := map[string]struct {
		atLeast  int
		atMost   int
		input    optional.String
		expected string
	}{
		"ok": {
			atLeast:  4,
			atMost:   6,
			input:    optional.StringFrom("exact"),
			expected: "",
		},
		"too long": {
			atLeast:  4,
			atMost:   6,
			input:    optional.StringFrom("too long"),
			expected: "must have at least 4 and at most 6 characters, but had 8",
		},
		"too short": {
			atLeast:  4,
			atMost:   6,
			input:    optional.StringFrom("sht"),
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
				Between: ok_range.IntBetween(c.atLeast, c.atMost),
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
