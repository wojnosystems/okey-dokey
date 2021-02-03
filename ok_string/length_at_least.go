package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultLengthAtLeastFormat(definition *LengthAtLeast, value optional.String) string {
	return fmt.Sprintf("cannot have fewer than %d characters", definition.Length)
}

type LengthAtLeast struct {
	Format func(definition *LengthAtLeast, value optional.String) string
	Length int
}

func (m *LengthAtLeast) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultLengthAtLeastFormat
	if m.Format != nil {
		formatter = m.Format
	}
	value.IfSet(func(actual string) {
		if len(actual) < m.Length {
			violationReceiver.Emit(formatter(m, value))
		}
	})
	return ok_action.Continue
}
