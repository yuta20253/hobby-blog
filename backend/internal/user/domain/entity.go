package user

type User struct {
	ID ID
	Name Name
	Email Email
	PasswordHash string
}

func NewUser(name Name, email Email, PasswordHash string) *User {
	return &User{
		Name: name,
		Email: email,
		PasswordHash: PasswordHash,
	}
}

func Reconstruct(id ID, name Name, email Email, PasswordHash string) *User {
	return &User{
		ID: id,
		Name: name,
		Email: email,
		PasswordHash: PasswordHash,
	}
}
