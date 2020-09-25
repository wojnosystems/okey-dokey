package ok_uint64

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"testing"
)

func TestEqual_Validate(t *testing.T) {
	cases := map[string]struct {
		eq       uint64
		input    optional.Uint64
		expected string
	}{
		"ok": {
			eq:       3,
			input:    optional.Uint64From(3),
			expected: "",
		},
		"too large": {
			eq:       6,
			input:    optional.Uint64From(7),
			expected: "must be exactly 6",
		},
		"equal": {
			eq:       5,
			input:    optional.Uint64From(4),
			expected: "must be exactly 5",
		},
		"nil": {
			eq: 5,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := Equal{
				Value: c.eq,
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
