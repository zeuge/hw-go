package entity

type Error string

const (
	ErrUnknownCommand Error = "unknown command"
)

func (e Error) Error() string {
	return string(e)
}
