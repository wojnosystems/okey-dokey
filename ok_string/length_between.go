package ok_string

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultLengthBetweenFormat(definition *LengthBetween, value optional.String) string {
	return fmt.Sprintf("must have at least %d and at most %d characters, but had %d", definition.AtLeast, definition.AtMost, len(value.Value()))
}

type LengthBetween struct {
	Format  func(definition *LengthBetween, value optional.String) string
	AtLeast int
	AtMost  int
}

func (m *LengthBetween) Validate(value optional.String, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultLengthBetweenFormat
	if m.Format != nil {
		formatter = m.Format
	}
	if !value.IsSet() {
		return ok_action.Continue
	}
	if len(value.Value()) < m.AtLeast || m.AtMost < len(value.Value()) {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
