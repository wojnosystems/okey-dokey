package ok_slice_int16

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultItemCountAtLeastFormat(definition *ItemCountAtLeast, value []optional.Int16) string {
	return fmt.Sprintf("cannot have fewer than %d items", definition.AtLeast)
}

type ItemCountAtLeast struct {
	Format  func(definition *ItemCountAtLeast, value []optional.Int16) string
	AtLeast int
}

func (m *ItemCountAtLeast) Validate(value []optional.Int16, violationReceiver bad.Emitter) ok_action.Enum {
	formatter := defaultItemCountAtLeastFormat
	if m.Format != nil {
		formatter = m.Format
	}
	length := 0
	if value != nil {
		length = len(value)
	}
	if length < m.AtLeast {
		violationReceiver.Emit(formatter(m, value))
	}
	return ok_action.Continue
}
