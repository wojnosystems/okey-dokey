package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"testing"
)

func TestLengthExactly_Validate(t *testing.T) {
	cases := map[string]struct {
		exactly  int
		input    optional.String
		expected string
	}{
		"ok": {
			exactly:  5,
			input:    optional.StringFrom("exact"),
			expected: "",
		},
		"too long": {
			exactly:  5,
			input:    optional.StringFrom("in-exact"),
			expected: "was not exactly 5 characters",
		},
		"nil": {
			exactly: 5,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := LengthExactly{
				Length: c.exactly,
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
