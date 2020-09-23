package ok_slice

import "github.com/wojnosystems/go-optional"

type optionalSliceString []optional.String

func (s *optionalSliceString) Len() int {
	if s == nil {
		return 0
	}
	return len(*s)
}
