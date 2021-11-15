package postgres

import (
	"database/sql"

	"github.com/RomaBiliak/lets-go-chat/internal/models"
)

func NewPostgreUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) GetUserByName(name string) (models.User, error) {
	user := models.User{}

	err := r.db.QueryRow("SELECT * FROM users WHERE name=$1", name).Scan(&user.Id, &user.Name, &user.Password)

	if err != nil && err != sql.ErrNoRows {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) CheckUserExists(name string) (bool, error) {
	id := 0

	err := r.db.QueryRow("SELECT id FROM users WHERE name=$1", name).Scan(&id)

	if err != nil && err != sql.ErrNoRows  {
		return false, err
	}

	if id > 0 {
		return true, nil
	}

	return false, nil
}

func (r *UserRepository) GetUserById(id models.UserId) (models.User, error) {
	user := models.User{}

	err := r.db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Name, &user.Password)

	if err != nil && err != sql.ErrNoRows {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user models.User) (models.UserId, error) {
	id := 0
	err := r.db.QueryRow("INSERT INTO users (name, password) VALUES ($1, $2)  RETURNING id", user.Name, user.Password).Scan(&id)

	if err != nil {
		return 0, err
	}

	return models.UserId(id), nil
}
