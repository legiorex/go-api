package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Token string `json:"token"`
}
