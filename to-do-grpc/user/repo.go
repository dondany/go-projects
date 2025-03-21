package user

import (
	"database/sql"
	"time"
)

type UserRepository interface {
	CreateUser(User) (User, error)
	GetUserByEmail(string) (User, error)
}

type PostgreUserRepository struct {
	db *sql.DB
}

func NewPostgreUserRepository(db *sql.DB) PostgreUserRepository {
	return PostgreUserRepository{
		db: db,
	}
}

func (r *PostgreUserRepository) CreateUser(user User) (User, error) {
	now := time.Now()
	query := "insert into users (name, email, password, created_at) values ($1, $2, $3, $4) returning id, name, email, created_at"
	var newUser User
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, now).Scan(&newUser.ID, &newUser.Name, &newUser.Email, &newUser.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return newUser, nil
}

func (r *PostgreUserRepository) GetUserByEmail(email string) (User, error) {
	query := "select id, name, email, password from users where email = $1"

	var user User
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}