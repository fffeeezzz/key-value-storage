package service

import (
	"encoding/gob"
	"fmt"
	"os"

	"key-value-storage-v1/internal/domain"
)

var datafile = "./files/tmp/datafile.gob"

func Load() error {
	fmt.Println("Loading", datafile)
	loadFrom, err := os.Open(datafile)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer loadFrom.Close()

	decoder := gob.NewDecoder(loadFrom)
	if err = decoder.Decode(&domain.Data); err != nil {
		return fmt.Errorf("decode: %w", err)
	}

	return nil
}
