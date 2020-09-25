package ok_int16

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"testing"
)

func TestIsRequired_Validate(t *testing.T) {
	cases := map[string]struct {
		input    optional.Int16
		expected string
	}{
		"ok": {
			input:    optional.Int16From(0),
			expected: "",
		},
		"missing": {
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
