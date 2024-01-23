package data

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) *User {
	return &User{
		0,
		name,
		email,
		password,
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
