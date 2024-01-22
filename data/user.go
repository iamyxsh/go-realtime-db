package data

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) *User {
	return &User{
		name,
		email,
		password,
	}
}

func (u *User) CreateUser() error {
	_, err := DB.Exec("INSERT INTO user (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)

	return err
}
