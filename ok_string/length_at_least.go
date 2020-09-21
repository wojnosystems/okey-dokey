package ok_string

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLengthAtLeastFormat(definition *LengthAtLeast, value *string) string {
	return fmt.Sprintf("cannot have fewer than %d characters", definition.Length)
}

type LengthAtLeast struct {
	Format func(definition *LengthAtLeast, value *string) string
	Length int
}

func (m *LengthAtLeast) Validate(value *string, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultLengthAtLeastFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if len(*value) < m.Length {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
