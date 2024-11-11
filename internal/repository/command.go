package repository

import (
	"context"

	"github.com/Kridalll/Bashhelper/internal/entity"
	"github.com/Kridalll/Bashhelper/pkg/postgres"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type commandRepository struct {
	db  *postgres.Postgres
	rdb *redis.Client
	// как долно хранить pid исполняемой команды (в секундах)
	pidTTL int64
	log    *logrus.Logger
}

func NewCommandRepository(pg *postgres.Postgres, redis *redis.Client, pidTTL int64, logger *logrus.Logger) *commandRepository {
	return &commandRepository{
		db:     pg,
		rdb:    redis,
		pidTTL: pidTTL,
		log:    logger,
	}
}

func (r *commandRepository) CreateCommand(ctx context.Context, commandText string) (entity.Command, error) {
	return entity.Command{}, nil
}
func (r *commandRepository) DeleteCommandById(ctx context.Context, commandId uint64) error {
	return nil
}

func (r *commandRepository) ListCommands(ctx context.Context, limit, offset uint64) ([]entity.Command, error) {
	return nil, nil
}

func (r *commandRepository) GetCommandById(ctx context.Context, commandId uint64) (entity.Command, error) {
	return entity.Command{}, ErrCommandNotFound
}

func (r *commandRepository) SaveCommandOutput(ctx context.Context, commandId uint64, line string) error {
	return nil
}
func (r *commandRepository) ClearCommandOutput(ctx context.Context, commandId uint64) error {
	return nil
}

// тут будет кэшироваться pid команды
func (r *commandRepository) SetCommandPID(ctx context.Context, commandId uint64, pid int) error {
	return nil
}

// кэш pid команды
func (r *commandRepository) GetCommandPID(ctx context.Context, commandId uint64) (int, error) {
	return 0, ErrCommandNotRunning
}

// кэш pid команды
func (r *commandRepository) DeleteCommandPID(ctx context.Context, commandId uint64) error {
	return nil
}
