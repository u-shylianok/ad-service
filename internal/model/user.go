package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HashPassword(password string) (string, error) {
	const COST int = 10 // just for example (no requirements are provided)

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), COST) // $2a$10$TVXESRL0IiN3YdNJjKAuDe4j8K3ggwPfqoUsQz73YWPiwHikyhJxG
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
