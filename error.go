package hydros

// ErrorResponse error response payload
type ErrorResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}
