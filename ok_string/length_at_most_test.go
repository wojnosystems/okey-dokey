package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"testing"
)

func TestLengthAtMost_Validate(t *testing.T) {
	cases := map[string]struct {
		atMost   int
		input    optional.String
		expected string
	}{
		"ok": {
			atMost:   20,
			input:    optional.StringFrom("short"),
			expected: "",
		},
		"too long": {
			atMost:   5,
			input:    optional.StringFrom("lessThan20"),
			expected: "cannot have more than 5 characters",
		},
		"nil": {
			atMost: 5,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := LengthAtMost{
				Length: c.atMost,
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
