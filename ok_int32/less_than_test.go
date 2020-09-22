package ok_int32

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"testing"
)

func TestLessThan_Validate(t *testing.T) {
	cases := map[string]struct {
		lessThan int32
		input    optional.Int32
		expected string
	}{
		"ok": {
			lessThan: 3,
			input:    optional.Int32From(2),
			expected: "",
		},
		"too large": {
			lessThan: 6,
			input:    optional.Int32From(7),
			expected: "must be less than 6",
		},
		"equal": {
			lessThan: 5,
			input:    optional.Int32From(5),
			expected: "must be less than 5",
		},
		"nil": {
			lessThan: 5,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := LessThan{
				Value: c.lessThan,
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
