package request

type UserRequest struct {
	Email string `json:"email" binding:""`
	Password string `json:"password" binding:""`
	Name string `json:"name" binding:"required,min=4,max=100"`
	Age int8 `json:"age" binding:""`
}