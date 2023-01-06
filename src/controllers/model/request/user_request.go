package request

type UserRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	name string `json:"name"`
	age int8 `json:"age"`
}