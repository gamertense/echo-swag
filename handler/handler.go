package handler

// Controller example
type Handler struct {
}

// NewController example
func NewHandler() *Handler {
	return &Handler{}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}
