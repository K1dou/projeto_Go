package response

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required,min=3,max=50"`
	Age   int    `json:"age" binding:"required,min=18,max=120"`
}
