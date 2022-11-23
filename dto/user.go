package dto

type UserCreate struct {
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
	Email     string `json:"email" form:"email" binding:"required" validate:"email" default:"user"`
	Password  string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
}
type UserUpdate struct {
	ID        uint64 `json:"id,string,omitempty" form:"id" binding:"required"`
	Password  string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
}
