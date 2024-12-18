package service

import (
	"context"

	"github.com/Kridalll/Bashhelper/internal/entity"
)

type CommandService interface {
	CreateCommand(ctx context.Context, commandText string) (entity.Command, error)
	DeleteCommandById(ctx context.Context, commandId uint64) error
	ListCommands(ctx context.Context, limit, offset uint64) ([]entity.Command, error)
	GetCommandById(ctx context.Context, commandId uint64) (entity.Command, error)
	RunCommand(ctx context.Context, commandId uint64) error
	StopCommand(ctx context.Context, commandId uint64) error
}

var _ CommandService = &commandService{}
