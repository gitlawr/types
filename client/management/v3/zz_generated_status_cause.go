package client

const (
	StatusCauseType         = "statusCause"
	StatusCauseFieldField   = "field"
	StatusCauseFieldMessage = "message"
	StatusCauseFieldType    = "reason"
)

type StatusCause struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"reason,omitempty"`
}
