package repository

import (
	"context"

	"github.com/Kridalll/Bashhelper/internal/entity"
)

type CommandRepository interface {
	CreateCommand(ctx context.Context, commandText string) (entity.Command, error)
	DeleteCommandById(ctx context.Context, commandId uint64) error
	ListCommands(ctx context.Context, limit, offset uint64) ([]entity.Command, error)
	GetCommandById(ctx context.Context, commandId uint64) (entity.Command, error)
	SaveCommandOutput(ctx context.Context, commandId uint64, line string) error
	ClearCommandOutput(ctx context.Context, commandId uint64) error
	SetCommandPID(ctx context.Context, commandId uint64, pid int) error
	GetCommandPID(ctx context.Context, commandId uint64) (int, error)
	DeleteCommandPID(ctx context.Context, commandId uint64) error
}

var _ CommandRepository = &commandRepository{}
