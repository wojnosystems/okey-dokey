package ok_string

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLengthAtMostFormat(definition *LengthAtMost, value *string) string {
	return fmt.Sprintf("cannot have more than %d characters", definition.Length)
}

type LengthAtMost struct {
	Format func(definition *LengthAtMost, value *string) string
	Length int
}

func (m *LengthAtMost) Validate(value *string, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultLengthAtMostFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if len(*value) > m.Length {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
