package ok_slice_int

import "okey-dokey/ok_int"

type On struct {
	Id          string
	Ensure      []Definer
	EnsureItems []ok_int.Definer
}
