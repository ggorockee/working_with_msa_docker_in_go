package helpers

type RegisterUserPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type CreateMemoPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateMemoPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPayload struct {
	Name string `json:"name"`
}