package ok_int8

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"testing"
)

func TestLessThanOrEqual_Validate(t *testing.T) {
	cases := map[string]struct {
		lessThanOrEq int8
		input        optional.Int8
		expected     string
	}{
		"ok": {
			lessThanOrEq: 3,
			input:        optional.Int8From(2),
			expected:     "",
		},
		"too large": {
			lessThanOrEq: 6,
			input:        optional.Int8From(7),
			expected:     "must be less than or equal to 6",
		},
		"equal": {
			lessThanOrEq: 5,
			input:        optional.Int8From(5),
			expected:     "",
		},
		"nil": {
			lessThanOrEq: 5,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := LessThanOrEqual{
				Value: c.lessThanOrEq,
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
