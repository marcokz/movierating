package entity

type Rating struct {
	User  User
	Movie Movie
	Ratio uint8
}
