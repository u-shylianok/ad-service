package secure

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type Secure struct {
	Hasher
}

//counterfeiter:generate --fake-name HasherMock -o ../../testing/mocks/secure/hash.go . Hasher
type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

func NewSecure() *Secure {
	return &Secure{
		Hasher: NewHash(),
	}
}
