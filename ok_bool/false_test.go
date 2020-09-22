package ok_bool

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"testing"
)

func TestFalse_Validate(t *testing.T) {
	cases := map[string]struct {
		input    optional.Bool
		expected string
	}{
		"ok": {
			input:    optional.BoolFrom(false),
			expected: "",
		},
		"not ok": {
			input:    optional.BoolFrom(true),
			expected: "must be false",
		},
		"nil": {},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := False{}
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
