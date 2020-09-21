package ok_int64

import (
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultIsRequiredFormat(definition *IsRequired, value *int64) string {
	return "is required"
}

type IsRequired struct {
	Format func(definition *IsRequired, value *int64) string
}

func (m *IsRequired) Validate(value *int64, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultIsRequiredFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if value == nil {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
