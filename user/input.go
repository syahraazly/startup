package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required" form:"name"`
	Occupation string `json:"occupation" binding:"required" form:"occupation"`
	Email      string `json:"email" binding:"required,email" form:"email"`
	Password   string `json:"password" binding:"required" form:"password"`
}
