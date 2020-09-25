package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLengthAtMostFormat(definition *LengthAtMost, value optional.String) string {
	return fmt.Sprintf("cannot have more than %d characters", definition.Length)
}

type LengthAtMost struct {
	Format func(definition *LengthAtMost, value optional.String) string
	Length int
}

func (m *LengthAtMost) Validate(value optional.String, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultLengthAtMostFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if len(value.Value()) > m.Length {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
