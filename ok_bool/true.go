package ok_bool

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultTrueFormat(definition *True, value optional.Bool) string {
	return "must be true"
}

type True struct {
	Format func(definition *True, value optional.Bool) string
}

func (m *True) Validate(value optional.Bool, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultTrueFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if !value.Value() {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
