package models

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
