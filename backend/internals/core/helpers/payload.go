package helpers

// Common
type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Memo
type CreateMemoPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateMemoPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// User
type RegisterUserPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPayload struct {
	Name string `json:"name"`
}
