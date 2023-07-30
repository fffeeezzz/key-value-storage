package service

import (
	"encoding/gob"
	"fmt"
	"os"

	"key-value-storage-v1/internal/domain"
)

func Save() error {
	fmt.Println("Saving", datafile)
	err := os.Remove(datafile)
	if err != nil {
		return fmt.Errorf("remove: %w", err)
	}

	saveTo, err := os.Create(datafile)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(domain.Data)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}
