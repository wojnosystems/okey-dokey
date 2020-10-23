package ok_string

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"regexp"
	"testing"
)

func TestMatchesRegexp_Validate(t *testing.T) {
	cases := map[string]struct {
		input    optional.String
		expected string
	}{
		"ok": {
			input:    optional.StringFrom("127.0.0.1"),
			expected: "",
		},
		"not match": {
			input:    optional.StringFrom("notipaddress"),
			expected: "did not match the pattern",
		},
		"nil": {
			expected: "",
		},
	}

	pattern := regexp.MustCompile(`^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$`)

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			validator := MatchesRegexp{
				Pattern: pattern,
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
