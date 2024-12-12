package models

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type SigninRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Auth0Response struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
}
