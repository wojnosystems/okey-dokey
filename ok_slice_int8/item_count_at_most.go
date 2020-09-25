package ok_slice_int8

import (
	"fmt"
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
)

func defaultItemCountAtMostFormat(definition *ItemCountAtMost, value []optional.Int8) string {
	return fmt.Sprintf("cannot have more than %d items", definition.AtMost)
}

type ItemCountAtMost struct {
	Format func(definition *ItemCountAtMost, value []optional.Int8) string
	AtMost int
}

func (m *ItemCountAtMost) Validate(value []optional.Int8, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultItemCountAtMostFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = len(value)
	}
	if length > m.AtMost {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
