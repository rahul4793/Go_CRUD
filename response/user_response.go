package response

type UserResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
