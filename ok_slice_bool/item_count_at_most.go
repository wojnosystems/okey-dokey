package ok_slice_bool

import (
	"fmt"
	"github.com/wojnosystems/go-optional/v2"
	"github.com/wojnosystems/okey-dokey/bad"
	"github.com/wojnosystems/okey-dokey/ok_action"
)

func defaultItemCountAtMostFormat(definition *ItemCountAtMost, value []optional.Bool) string {
	return fmt.Sprintf("cannot have more than %d items", definition.AtMost)
}

type ItemCountAtMost struct {
	Format func(definition *ItemCountAtMost, value []optional.Bool) string
	AtMost int
}

func (m *ItemCountAtMost) Validate(value []optional.Bool, violationReceiver bad.Emitter) ok_action.Enum {
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
