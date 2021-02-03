package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"testing"
)

func TestIsURL_Validate(t *testing.T) {
	cases := map[string]struct {
		input    optional.String
		expected string
	}{
		"ok": {
			input:    optional.StringFrom("https://www.example.com"),
			expected: "",
		},
		"not URL": {
			input:    optional.StringFrom(":not url"),
			expected: "is not a URL",
		},
		"nil": {
			expected: "",
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := IsURL{}
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
