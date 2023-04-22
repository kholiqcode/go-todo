package dtos

type CreateActivityGroupRequest struct {
	Title string `json:"title" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
}
