package ok_slice_bool

import "okey-dokey/ok_bool"

type On struct {
	Id          string
	Ensure      []Definer
	EnsureItems []ok_bool.Definer
}
