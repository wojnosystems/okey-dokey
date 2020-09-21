package ok_string

import (
	"fmt"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultFormatLengthExactly(definition *LengthExactly, value *string) string {
	return fmt.Sprintf("was not exactly %d characters", definition.Length)
}

type LengthExactly struct {
	Format func(definition *LengthExactly, value *string) string
	Length int
}

func (m *LengthExactly) Validate(value *string, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultFormatLengthExactly
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		return ok_action.Continue
	}
	if len(*value) != m.Length {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
