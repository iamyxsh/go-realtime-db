package data

import "github.com/jmoiron/sqlx"

type AuthUser struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewAuthUser(name, email, password string) *AuthUser {
	return &AuthUser{
		0,
		name,
		email,
		password,
	}
}

func (u *AuthUser) CreateUser(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO users9999 (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)

	return err
}

func (u *AuthUser) GetUserByEmail(db *sqlx.DB) error {
	err := db.Get(u, "SELECT * FROM users9999 WHERE email = $1", u.Email)
	return err
}

func (u *AuthUser) GetUserById(db *sqlx.DB) error {
	err := db.Get(u, "SELECT * FROM users9999 WHERE id = $1", u.Id)
	return err
}

func (u *AuthUser) SaveUser(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE users9999 SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, u.Id)

	return err
}
