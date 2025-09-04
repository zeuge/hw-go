package entity

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrNotFound Error = "not found"
)
