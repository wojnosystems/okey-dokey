package ok_slice_uint8

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultItemCountAtLeastFormat(definition *ItemCountAtLeast, value []optional.Uint8) string {
	return fmt.Sprintf("cannot have fewer than %d items", definition.AtLeast)
}

type ItemCountAtLeast struct {
	Format  func(definition *ItemCountAtLeast, value []optional.Uint8) string
	AtLeast int
}

func (m *ItemCountAtLeast) Validate(value []optional.Uint8, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultItemCountAtLeastFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = len(value)
	}
	if length < m.AtLeast {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
