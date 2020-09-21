package ok_uint8

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultIsRequiredFormat(definition *IsRequired, value *uint8) string {
	return "is required"
}

type IsRequired struct {
	Format func(definition *IsRequired, value *uint8) string
}

func (m *IsRequired) Validate(value *uint8, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultIsRequiredFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
