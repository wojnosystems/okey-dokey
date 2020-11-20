package bad

type Collector interface {
	// IsEmpty is true if Collector has no validation errors
	IsEmpty() bool
	// HasAny is true if there is at least 1 validation error
	HasAny() bool
	// Paths returns the list of paths with validation errors. Skips any paths without errors
	Paths() (paths []string)
	// MessagesAtPath returns the errors for a path or empty array if there are no errors or the path does not exist
	MessagesAtPath(path string) []string
}

type ReceiveCollector interface {
	MemberEmitter
	Collector
}
