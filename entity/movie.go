package entity

type Movie struct {
	ID          int64
	Title       string
	Year        int64
	Director    []Person
	Actors      []Person
	Description string
}
