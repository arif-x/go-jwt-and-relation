package models

type User struct {
	ID       int             `json:"id" form:"id" gorm:"primaryKey"`
	Name     string          `json:"name" form:"name" gorm:"not null"`
	Email    string          `json:"email" form:"email" gorm:"not null; index:unique"`
	Password string          `json:"-" form:"password" gorm:"not null"`
	Locker   *LockerResponse `json:"locker"`
	Posts    []PostResponse  `json:"posts"`
}

type UserStore struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (UserStore) TableName() string {
	return "users"
}

type UserResponse struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
