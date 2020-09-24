package ok_slice_int64

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
	"strconv"
)

func Validate(values []optional.Int64, on *On, violations bad.MemberReceiver) {
	mem := violations.MessageReceiver(on.Id)
	for _, definition := range on.Ensure {
		action := definition.Validate(values, mem)
		if action != ok_action.Continue {
			return
		}
	}
	for i, value := range values {
		mem := violations.MessageReceiver(on.Id + "[" + strconv.FormatInt(int64(i), 10) + "]")
		for _, definition := range on.EnsureItems {
			action := definition.Validate(value, mem)
			if action != ok_action.Continue {
				return
			}
		}
	}
}
