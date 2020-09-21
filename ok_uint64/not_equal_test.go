package ok_uint64

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"okey-dokey/bad"
	"testing"
)

func TestNotEqual_Validate(t *testing.T) {
	cases := map[string]struct {
		notEq    uint64
		input    *uint64
		expected string
	}{
		"ok less": {
			notEq:    3,
			input:    addrOf(2),
			expected: "",
		},
		"ok greater": {
			notEq:    3,
			input:    addrOf(4),
			expected: "",
		},
		"equal": {
			notEq:    3,
			input:    addrOf(3),
			expected: "must not be 3",
		},
		"nil": {
			notEq: 3,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := NotEqual{
				Value: c.notEq,
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
