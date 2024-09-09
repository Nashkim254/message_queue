package main

import "fmt"

type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

type MemoryStore struct {
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStore) Push(b []byte) (int, error) {
	s.data = append(s.data, b)
	return len(s.data) - 1, nil
}

func (s *MemoryStore) Fetch(offset int) ([]byte, error) {
	if len(s.data) < offset {
		return nil, fmt.Errorf("Offset (%d) is too high", offset)
	}
	return s.data[offset], nil
}
