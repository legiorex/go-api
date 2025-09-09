package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}
type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type LoginPayload struct {
	Token string `json:"token"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type RegisterPayload struct {
	Token string `json:"token"`
	Email string `json:"email"`
}
