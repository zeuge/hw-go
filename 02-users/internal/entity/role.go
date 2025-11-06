package entity

type Role string

func (r Role) String() string {
	return string(r)
}

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
	GuestRole Role = "guest"
)
