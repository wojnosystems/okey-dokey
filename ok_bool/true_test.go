package ok_bool

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"okey-dokey/bad"
	"testing"
)

func TestTrue_Validate(t *testing.T) {
	cases := map[string]struct {
		input    *bool
		expected string
	}{
		"ok": {
			input:    addrOf(true),
			expected: "",
		},
		"not ok": {
			input:    addrOf(false),
			expected: "must be true",
		},
		"nil": {},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := True{}
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
