package domain

type User struct {
	Name string
	Email string
	Password string
}

func CreateUser(name, email, password string) *User {
	if name == "" || email == "" || password == "" {
		return nil
	}

	return &User{
		Name: name,
		Email: email,
		Password: password,
	}
}
