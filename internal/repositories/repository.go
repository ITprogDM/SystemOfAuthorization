package repositories

import (
	"SystemOfAuthorization/internal/models"
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	UserRepository
	SessionRepository
	RoleRepository
}

type UserRepository interface {
	// Создание пользователя
	CreateUser(ctx context.Context, user *models.User) error

	//Получение пользователя
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	//Обновление данных пользователя
	UpdateUser(ctx context.Context, user *models.User) error
	UpdatePassword(ctx context.Context, id uuid.UUID, newHashedPassword string) error

	//Удаление пользователя
	DeleteUser(ctx context.Context, id uuid.UUID) error

	//Активация пользователя
	ActivateUser(ctx context.Context, id uuid.UUID) error
	DeactivateUser(ctx context.Context, id uuid.UUID) error
}

type SessionRepository interface {
	//Создание сессии
	CreateSession(ctx context.Context, id uuid.UUID) (*models.Session, error)

	//Получение сессиии
	GetSessionByID(ctx context.Context, id uuid.UUID) (*models.Session, error)
	GetSessionByToken(ctx context.Context, token string) (*models.Session, error)
	GetUserSessions(ctx context.Context, userID uuid.UUID) ([]*models.Session, error)

	//Удаление сессии
	DeleteSession(ctx context.Context, id uuid.UUID) error
	DeleteAllUserSessions(ctx context.Context, userID uuid.UUID) error

	//Проверка активности
	IsValidSession(ctx context.Context, token string) (bool, error)
}

type RoleRepository interface {
	// Если роли динамические
	AssignRole(ctx context.Context, userID uuid.UUID, roleID int) error
	RevokeRole(ctx context.Context, userID uuid.UUID, roleID int) error
	GetUserRoles(ctx context.Context, userID uuid.UUID) ([]*models.Role, error)

	// если роли статические
	IsAdmin(ctx context.Context, userID uuid.UUID) (bool, error)
}
