package database

import (
	"context"
	"movieraiting/entity"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase(pool *pgxpool.Pool) *Database {
	return &Database{pool: pool}
}

func (d *Database) SignUpDB(ctx context.Context, u entity.User) error {
	_, err := d.pool.Exec(ctx, "insert into users (login, email, password_hash) values ($1, $2, $3)", u.Login, u.Email, u.PasswordHash)
	return err
}

func (d *Database) CreateMovieDB(ctx context.Context, u entity.Movie) error { /// ??rest.CreateMovieRequest
	_, err := d.pool.Exec(ctx, "insert into movierating (??[]Person принимает же entity.Movie) values ($1, $2, $3, $4, $5)", u.Title, u.Year, u.Director, u.Actors, u.Description)
	return err
}

func (d *Database) GetMovie(ctx context.Context) ([]entity.Movie, error) {
	rows, err := d.pool.Query(ctx, "select * from movie") // для нескольких данных
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movie []entity.Movie
	for rows.Next() {
		var u entity.Movie
		err := rows.Scan(
			&u.ID,
			&u.Title,
			&u.Year,
			&u.Director,
			&u.Actors,
			&u.Description,
		)
		if err != nil {
			return nil, err
		}
		movie = append(movie, u)
	}
	return movie, nil
}
func (d *Database) GetMovieByTitle(ctx context.Context, title string) {
	row := d.pool.QueryRow(ctx, "select title from movierating where title=$1", title)
}

func (d *Database) GetMovieByDirecor(ctx context.Context, id int64) (entity.Person, error) { /// ?? праивльно?
	row := d.pool.QueryRow(ctx, "select id, firstname, lastname from person where id=$1", id)

	var director entity.Person
	err := row.Scan(&director.ID, &director.FirstName, &director.LastName)
	if err != nil {
		return entity.Person{}, err
	}
	return entity.Person{}, nil
}

func (d *Database) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	row := d.pool.QueryRow(ctx, "select id, login, email, password_hash from users where login=$1", login)

	var user entity.User
	err := row.Scan(&user.ID, &user.Login, &user.Email, &user.PasswordHash)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
