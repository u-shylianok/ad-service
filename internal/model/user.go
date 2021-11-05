package model

import "golang.org/x/crypto/bcrypt"

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name" binding:"required"`
	Username string `db:"username" binding:"required"`
	Password string `db:"password" binding:"required"`
}

type UserResponse struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		Name:     u.Name,
		Username: u.Username,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
