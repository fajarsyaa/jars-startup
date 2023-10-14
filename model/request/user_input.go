package request

type RegisterUserRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Occupation string `json:"occupation" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AvailableEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}
