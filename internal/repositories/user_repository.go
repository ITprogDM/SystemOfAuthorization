package repositories

import (
	"SystemOfAuthorization/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db  *pgxpool.Pool
	log *logrus.Logger
}

func NewUserRepository(db *pgxpool.Pool, log *logrus.Logger) *UserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) error {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		r.log.Error("Ошибка создания пользователя в базе данных")
		return err
	}

	r.log.Info("Пользователь успешно создан")
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User

	query := "SELECT username, email FROM users WHERE ID = $1"
	err := r.db.QueryRow(ctx, query, userID).Scan(&user.Username, &user.Email)
	if err != nil {
		r.log.Error("Пользователь не найден в базе данных по ID!")
		return nil, err
	}

	r.log.Infof("Пользователь с ID: %d найден!", userID)
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	query := "SELECT username, email FROM users WHERE username = $1"
	err := r.db.QueryRow(ctx, query, username).Scan(&user.Username, &user.Email)
	if err != nil {
		r.log.Error("Пользователь не найден в БД по username!")
		return nil, err
	}

	r.log.Infof("Пользователь с Username: %s найден!", username)
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := "SELECT username, email FROM users WHERE email = $1"
	err := r.db.QueryRow(ctx, query, email).Scan(&user.Username, &user.Email)
	if err != nil {
		r.log.Error("Пользователь не найден в БД по email!")
		return nil, err
	}

	r.log.Infof("Пользователь с Email: %s найден!", email)
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user models.User) error {
	// Todo
}
