package entity

type User struct {
	ID           int
	Login        string
	Email        string
	PasswordHash []byte
}
