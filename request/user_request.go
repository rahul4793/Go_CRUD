package request

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=2"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18,lte=80"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"omitempty,min=2"`
	Email string `json:"email" validate:"omitempty,email"`
	Age   int    `json:"age" validate:"omitempty,gte=18,lte=80"`
}
