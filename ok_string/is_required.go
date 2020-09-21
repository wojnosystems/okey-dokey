package ok_string

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultIsRequiredFormat(definition *IsRequired, value *string) string {
	return "is required"
}

type IsRequired struct {
	Format func(definition *IsRequired, value *string) string
}

func (m *IsRequired) Validate(value *string, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultIsRequiredFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
