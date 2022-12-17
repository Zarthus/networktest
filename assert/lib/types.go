package lib

type Args struct {
	V4             bool
	V6             bool
	TimeoutSeconds int
	Host           string
}

type Status int

const (
	Ok          Status = 0
	Warn        Status = 1
	Error       Status = 2
	NotExecuted Status = 3
)

type Result struct {
	Status    Status
	Check     string
	Assertion string
	Guidance  []string
}
