package internal

import (
	"fmt"

	"key-value-storage/internal/domain"
	"key-value-storage/internal/service"
)

type InsertCommand struct {
	Key     string
	ID      string
	Name    string
	Surname string
}

type InsertUseCase struct{}

func (useCase *InsertUseCase) Handle(command *InsertCommand) error {
	value := &domain.Element{
		ID:      command.ID,
		Name:    command.Name,
		Surname: command.Surname,
	}

	if !domain.Data.Add(command.Key, value) {
		return fmt.Errorf("add operation failed")
	}

	if err := service.Save(); err != nil {
		return fmt.Errorf("save: %w", err)
	}

	return nil
}
