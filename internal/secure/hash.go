package secure

import "golang.org/x/crypto/bcrypt"

type Hash struct {
	cost int
}

func NewHash() *Hash {
	return &Hash{cost: bcrypt.DefaultCost}
}

func (h *Hash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

func (h *Hash) CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
