package data

import "github.com/jmoiron/sqlx"

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	APIKey   string `json:"apiKey"`
}

func NewUser(name, email, password string) *User {
	return &User{
		0,
		name,
		email,
		password,
		"",
	}
}

func (u *User) CreateUser(db *sqlx.DB) error {
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)

	return err
}

func (u *User) GetUserByEmail(db *sqlx.DB) error {
	err := db.Get(u, "SELECT * FROM users WHERE email = $1", u.Email)
	return err
}

func (u *User) GetUserByApiKey(db *sqlx.DB) error {
	err := db.Get(u, "SELECT * FROM users WHERE apikey = $1", u.APIKey)

	return err
}

func (u *User) SaveUser(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE users SET name = $1, email = $2, apikey = $3 WHERE id = $4", u.Name, u.Email, u.APIKey, u.Id)

	return err
}
