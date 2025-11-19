package entity

type Error string

const (
	ErrNotFound Error = "not found"
)

func (e Error) Error() string {
	return string(e)
}
