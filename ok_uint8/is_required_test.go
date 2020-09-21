package ok_uint8

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"okey-dokey/bad"
	"testing"
)

func TestIsRequired_Validate(t *testing.T) {
	cases := map[string]struct {
		input    *uint8
		expected string
	}{
		"ok": {
			input:    addrOf(0),
			expected: "",
		},
		"missing": {
			input:    nil,
			expected: "is required",
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := IsRequired{}
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
