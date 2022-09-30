package user

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
