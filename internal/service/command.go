package service

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/Kridalll/Bashhelper/internal/entity"
	"github.com/Kridalll/Bashhelper/internal/repository"
	"github.com/sirupsen/logrus"
)

type commandService struct {
	commandRepo repository.CommandRepository
	// пусть до стандартной оболочки, например /bin/sh или /bin/bash
	defaultShellPath string
	log              *logrus.Logger
}

func NewCommandService(commandRepo repository.CommandRepository, defaultShellPath string, logger *logrus.Logger) *commandService {
	return &commandService{
		commandRepo:      commandRepo,
		defaultShellPath: defaultShellPath,
		log:              logger,
	}
}

func (s *commandService) CreateCommand(ctx context.Context, commandText string) (entity.Command, error) {
	// запрос к репо
	return entity.Command{}, nil
}

func (s *commandService) DeleteCommandById(ctx context.Context, commandId uint64) error {
	// запрос к репо
	return nil
}

func (s *commandService) ListCommands(ctx context.Context, limit, offset uint64) ([]entity.Command, error) {
	// запрос к репо
	return nil, nil
}

func (s *commandService) GetCommandById(ctx context.Context, commandId uint64) (entity.Command, error) {
	// запрос к репо
	return entity.Command{}, nil
}

func (s *commandService) processOutput(ctx context.Context, commandId uint64, line string) error {
	// запрос к репо
	return nil
}

func (s *commandService) RunCommand(ctx context.Context, commandId uint64) error {
	var command entity.Command
	// запущена ли уже команда?
	// запрос к репо
	// ... проверка - если да, возвращаем ErrCommandAlreadyRunning

	// получаем команду
	// запрос к репо

	// отчищаем старый вывод команды
	// запрос к репо

	cmd := exec.Command(s.defaultShellPath, "-c", command.Text)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// поскольку ставим false, сигналы к программе будут дублироваться на дочерние процессы
		// НО зато будут права на то, чтобы эти дочерние процессы останавливать
		Setpgid: false,
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		s.log.Errorf("commandService.RunCommand -> cmd.StdoutPipe: %v", err)
		return err
	}
	scanner := bufio.NewScanner(stdout)

	// стартуем исполнение команды
	err = cmd.Start()
	if err != nil {
		fmt.Println(cmd.Path, cmd.Args)
		s.log.Errorf("commandService.RunCommand -> cmd.Start: %v", err)
		return err
	}

	// сохраняем  pid команды
	// запрос к репо

	// при завершении команды или возникновении ошибки "забываем" её pid
	defer func() {
		// запрос к репо на удаление команды из кэша
	}()

	// сканируем и сохраняем вывод команды
	for scanner.Scan() {
		_ = s.processOutput(ctx, commandId, scanner.Text())
	}

	if scanner.Err() != nil {
		cmd.Process.Kill()
		cmd.Wait()
		s.log.Errorf("commandService.RunCommand -> scanner.Err: %v", err)
		return scanner.Err()
	}

	return cmd.Wait()
}

func (s *commandService) StopCommand(ctx context.Context, commandId uint64) error {
	// проверяем наличие команды
	// запрос к репо

	// получаем pid команды
	// запрос к репо

	// после завершения должны удалить pid команды
	defer func() {
		// запрос к репо на удаление команды из кэша
	}()

	var pid int
	var err error

	// нет ошибки на unix системах
	process, _ := os.FindProcess(pid)
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		// процесс уже был завершен
		s.log.Errorf("commandService.StopCommand -> process.Signal: %v", err)
		return nil
	}

	// попробуем завершить по хорошему :)
	err = process.Signal(syscall.SIGTERM)
	if err != nil {
		// по хорошему не хочет, значит будет по плохому >:(
		err = process.Signal(syscall.SIGKILL)
		if err != nil {
			s.log.Errorf("commandService.StopCommand -> process.Signal: %v", err)
			return err
		}
		return nil
	}

	// проверяем что процесс завершился
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		err = process.Signal(syscall.Signal(0))
		if err != nil {
			// процесс завершился
			return nil
		}
	}

	// по хорошему он снова не захотел >:(
	err = process.Signal(syscall.SIGKILL)
	if err != nil {
		s.log.Errorf("commandService.StopCommand -> process.Signal: %v", err)
		return err
	}
	return nil
}
