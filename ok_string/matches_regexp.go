package ok_string

import (
	"github.com/wojnosystems/go-optional"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"regexp"
)

func defaultMatchesRegexp(definition *MatchesRegexp, value optional.String) string {
	return "did not match the pattern"
}

type MatchesRegexp struct {
	Format  func(definition *MatchesRegexp, value optional.String) string
	Pattern *regexp.Regexp
}

func (m *MatchesRegexp) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultMatchesRegexp
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if !m.Pattern.MatchString(value.Value()) {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
