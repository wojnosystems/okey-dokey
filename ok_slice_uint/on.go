package ok_slice_uint

import "okey-dokey/ok_uint"

type On struct {
	Id          string
	Ensure      []Definer
	EnsureItems []ok_uint.Definer
}
