package admin

import (
	"context"
	"fmt"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/out"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserAdapter struct {
	db *gorm.DB
}

func (u *UserAdapter) CreateUser(ctx context.Context, user domain.User) error {
	var count int64

	_ = u.db.Find(&user, "username = ?", user.Username).Count(&count)
	if count != 0 {
		return fmt.Errorf("username %s already exists", user.Username)

	}
	result := u.db.Create(user)
	return result.Error
}

func (u *UserAdapter) User(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	result := u.db.Find(&user, "username = ?", username)
	return user, result.Error
}

func (u *UserAdapter) Users(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	result := u.db.Find(&users)
	return users, result.Error
}

func NewUserAdapter(config domain.PostgresConfig) out.UserPort {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.Database),
		), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	// db.AutoMigrate(&domain.DatabaseConnectConfig{})
	db.AutoMigrate(&domain.User{})
	return &UserAdapter{
		db: db,
	}
}
