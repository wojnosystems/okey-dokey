package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
	"github.com/wojnosystems/okey-dokey/ok_range"
)

func defaultLengthBetweenFormat(definition *LengthBetween, value string) string {
	return fmt.Sprintf("must have at least %d and at most %d characters, but had %d", definition.Between.Start(), definition.Between.End(), len(value))
}

type LengthBetween struct {
	Format  func(definition *LengthBetween, value string) string
	Between ok_range.Int
}

func (m *LengthBetween) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultLengthBetweenFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual string) {
		if len(actual) < m.Between.Start() || m.Between.End() < len(actual) {
			violationReceiver.Emit(formatter(m, actual))
		}
	})
	return ok_action.Continue
}
