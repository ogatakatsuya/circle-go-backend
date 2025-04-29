package schema

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token  string `json:"access_token"`
	UserID string `json:"user_id"`
}

type SignUpResponse struct {
	Token string `json:"access_token"`
}
