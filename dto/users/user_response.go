package usersdto

type UserResponse struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
