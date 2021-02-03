package ok_string

import (
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"regexp"
)

func defaultMatchesRegexp(definition *MatchesRegexp, value string) string {
	return "did not match the pattern"
}

type MatchesRegexp struct {
	Format  func(definition *MatchesRegexp, value string) string
	Pattern *regexp.Regexp
}

func (m *MatchesRegexp) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultMatchesRegexp
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual string) {
		if !m.Pattern.MatchString(actual) {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
