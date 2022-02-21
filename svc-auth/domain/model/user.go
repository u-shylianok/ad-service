package model

type UserRequest struct {
	Name     string
	Username string
	Password string
}

type SignInRequest struct {
	Username string
	Password string
}

type User struct {
	ID       uint32 `db:"id"`
	Name     string `db:"name" binding:"required"`
	Username string `db:"username" binding:"required"`
	Password string `db:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint32
	Name     string
	Username string
}

func (u User) ToResponse() UserResponse {
	return UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
	}
}
