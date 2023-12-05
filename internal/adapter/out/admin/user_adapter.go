package admin

import (
	"context"
	"fmt"
	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/out"
	"github.com/CorrectRoadH/Likit/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type UserAdapter struct {
	db *gorm.DB
}

func (u *UserAdapter) CreateUser(ctx context.Context, username string, password string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserAdapter) CreateUserByEnv() (domain.User, error) {
	// the first start create user
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	var user domain.User
	var count int64

	result := u.db.Find(&user, "username = ?", username).Count(&count)
	userData := &domain.User{
		Id:       utils.Uuid(),
		Username: username,
		Password: password,
	}

	// database not found username
	if count < 1 {
		// env not found username
		if len(username) < 1 {
			userData.Username = "likit"

		}
		// env not found password
		if len(password) < 1 {
			userData.Password = "likit"
		}
		result = u.db.Create(userData)
	}

	return user, result.Error
}

func NewUserAdapter(config domain.PostgresConfig) out.UserPortUseCase {
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
