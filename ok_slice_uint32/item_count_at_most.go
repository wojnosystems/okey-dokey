package ok_slice_uint32

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultItemCountAtMostFormat(definition *ItemCountAtMost, value []optional.Uint32) string {
	return fmt.Sprintf("cannot have more than %d items", definition.AtMost)
}

type ItemCountAtMost struct {
	Format func(definition *ItemCountAtMost, value []optional.Uint32) string
	AtMost int
}

func (m *ItemCountAtMost) Validate(value []optional.Uint32, violationReceiver bad.MessageReceiver) ok_action.Enum {
	formatter := defaultItemCountAtMostFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = len(value)
	}
	if length > m.AtMost {
		violationReceiver.ReceiveMessage(formatter(m, value))
	}
	return ok_action.Continue
}
