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

	_ = r.db.QueryRow("SELECT * FROM users WHERE name=$1", name).Scan(&user.Id, &user.Name, &user.Password)

	return user, nil
}

func (r *UserRepository) CreateUser(user models.User) (models.User, error) {
	_, err := r.db.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", user.Name, user.Password)

	if err != nil {
		return user, err
	}

	return r.GetUserByName(user.Name)
}
