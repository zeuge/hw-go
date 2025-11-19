package entity

type Error string

const (
	ErrNotFound             Error = "not found"
	ErrUnknownCommand       Error = "unknown command"
	ErrUnexpectedHTTPStatus Error = "unexpected HTTP status"
)

func (e Error) Error() string {
	return string(e)
}
