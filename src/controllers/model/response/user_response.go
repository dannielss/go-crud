package response

type UserResponse struct {
	ID string `json:"id"`
	Email string `json:"email"`
	name string `json:"name"`
	age int8 `json:"age"`
}