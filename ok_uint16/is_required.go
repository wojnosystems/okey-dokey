package ok_uint16

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultIsRequiredFormat(definition *IsRequired, value optional.Uint16) string {
	return "is required"
}

type IsRequired struct {
	Format func(definition *IsRequired, value optional.Uint16) string
}

func (m *IsRequired) Validate(value optional.Uint16, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultIsRequiredFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
