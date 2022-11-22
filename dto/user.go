package dto

type UserCreate struct {
	First_name string `json:"first_name" form:"first_name" binding:"required"`
	Last_name  string `json:"last_name" form:"last_name" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required" validate:"email"`
	Password   string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
}
