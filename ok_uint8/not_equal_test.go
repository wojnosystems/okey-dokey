package ok_uint8

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"testing"
)

func TestNotEqual_Validate(t *testing.T) {
	cases := map[string]struct {
		notEq    uint8
		input    optional.Uint8
		expected string
	}{
		"ok less": {
			notEq:    3,
			input:    optional.Uint8From(2),
			expected: "",
		},
		"ok greater": {
			notEq:    3,
			input:    optional.Uint8From(4),
			expected: "",
		},
		"equal": {
			notEq:    3,
			input:    optional.Uint8From(3),
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
