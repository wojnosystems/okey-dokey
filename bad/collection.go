package bad

import "strings"

// SliceMemberReceiver is a convenience implementation of MemberContainer used in tests, but applications should make their own
// so that they can handle notifications about errors
type collection struct {
	currentName string
	badFields   map[string][]string
}

func NewCollection() ReceiveCollector {
	return &collection{
		badFields: make(map[string][]string),
	}
}

func (r collection) IsEmpty() bool {
	return len(r.badFields) == 0
}

func (r collection) HasAny() bool {
	return !r.IsEmpty()
}

func (r collection) Paths() (paths []string) {
	paths = make([]string, 0, len(r.badFields))
	for path := range r.badFields {
		paths = append(paths, path)
	}
	return
}

func (r collection) MessagesAtPath(path string) (out []string) {
	out, _ = r.badFields[path]
	return
}

func (r *collection) Into(fieldName string) MemberEmitter {
	name := fieldName
	if r.currentName != "" {
		name = r.currentName
		if !strings.HasPrefix(fieldName, "[") {
			name += "."
		}
		name += fieldName
	}
	return &collection{
		currentName: name,
		badFields:   r.badFields,
	}
}

func (r *collection) Emit(message string) {
	r.badFields[r.currentName] = append(r.badFields[r.currentName], message)
}
