package data

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

func (u *User) CreateUser() error {
	_, err := DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)

	return err
}

func (u *User) GetUserByEmail() error {
	err := DB.Get(u, "SELECT * FROM users WHERE email = $1", u.Email)

	return err
}

func (u *User) GetUserByApiKey() error {
	err := DB.Get(u, "SELECT * FROM users WHERE apikey = $1", u.APIKey)

	return err
}

func (u *User) SaveUser() error {
	_, err := DB.Exec("UPDATE users SET name = $1, email = $2, apikey = $3 WHERE id = $4", u.Name, u.Email, u.APIKey, u.Id)

	return err
}
