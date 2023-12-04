package database

import (
	"context"
	"fmt"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/out"
	"github.com/jackc/pgx/v4"
)

type PostgresAdapter struct{}

func testPostgresConnection(ctx context.Context, conn *pgx.Conn) error {
	var greeting string
	err := conn.QueryRow(ctx, "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresAdapter) TestConnect(config domain.DatabaseConnectConfig) error {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	// 测试连接
	err = testPostgresConnection(ctx, conn)
	return err
}

func NewPostgresAdapter() out.PostgresUseCase {
	return &PostgresAdapter{}
}
