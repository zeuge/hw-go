package entity

type Action string

const (
	CreateAction Action = "create"
	GetAction    Action = "get"
	GetAllAction Action = "get_all"
	DeleteAction Action = "delete"
)

type Command struct {
	Action Action         `json:"action"`
	Data   map[string]any `json:"data,omitempty"`
}
