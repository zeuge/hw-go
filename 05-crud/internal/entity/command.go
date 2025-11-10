package entity

type Action string

const (
	ActionCreate Action = "create"
	ActionGet    Action = "get"
	ActionGetAll Action = "get_all"
	ActionDelete Action = "delete"
)

type Command struct {
	Action Action         `json:"action"`
	Data   map[string]any `json:"data,omitempty"`
}
