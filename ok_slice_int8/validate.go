package ok_slice_int8

import (
	"github.com/wojnosystems/go-optional"
	"okey-dokey/bad"
	"okey-dokey/ok_action"
	"strconv"
)

func Validate(values []optional.Int8, on *On, violations bad.MemberReceiver) {
	for _, definition := range on.Ensure {
		action := definition.Validate(values, violations)
		if action != ok_action.Continue {
			return
		}
	}
	for i, value := range values {
		mem := violations.MemberReceiver("[" + strconv.FormatInt(int64(i), 10) + "]")
		for _, definition := range on.EnsureItems {
			action := definition.Validate(value, mem)
			if action != ok_action.Continue {
				return
			}
		}
	}
}
