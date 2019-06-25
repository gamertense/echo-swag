package handler

// Handler example
type Handler struct {
}

// NewHandler example
func NewHandler() *Handler {
	return &Handler{}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}
