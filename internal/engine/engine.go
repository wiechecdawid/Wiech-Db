package engine

import (
	"errors"
)

type MemoryDB struct {
	data map[string]string
}

func (MemoryDB) New() *MemoryDB {
	return &MemoryDB{
		data: make(map[string]string),
	}
}

func (db *MemoryDB) Set(key, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	db.data[key] = value
	return nil
}

func (db *MemoryDB) Get(key string) (string, error){
	value, exists := db.data[key]
	if exists {
		return value, nil
	} 

	return "", errors.New("key not found")
}