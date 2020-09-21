package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"okey-dokey/bad"
	"testing"
)

func TestLengthAtLeast_Validate(t *testing.T) {
	cases := map[string]struct {
		atLeast  int
		input    *string
		expected string
	}{
		"ok": {
			atLeast:  3,
			input:    addrOf("short"),
			expected: "",
		},
		"too short": {
			atLeast:  6,
			input:    addrOf("short"),
			expected: "cannot have fewer than 6 characters",
		},
		"nil": {
			atLeast: 6,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := LengthAtLeast{
				Length: c.atLeast,
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
