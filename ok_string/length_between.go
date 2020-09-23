package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
	"okey-dokey/ok_range"
)

func defaultLengthBetweenFormat(definition *LengthBetween, value optional.String) string {
	return fmt.Sprintf("must have at least %d and at most %d characters, but had %d", definition.Between.Start(), definition.Between.End(), len(value.Value()))
}

type LengthBetween struct {
	Format  func(definition *LengthBetween, value optional.String) string
	Between ok_range.Int
}

func (m *LengthBetween) Validate(value optional.String, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultLengthBetweenFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if len(value.Value()) < m.Between.Start() || m.Between.End() < len(value.Value()) {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
